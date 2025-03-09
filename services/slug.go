package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/tools"
)

type slugUpdateSignal struct {
	Db        dbx.Builder
	SlugGroup string
	Err       chan error
}

var slugUpdateChannel = make(chan slugUpdateSignal)

func startUpdateSlugService(app *pocketbase.PocketBase) error {
	go func() {
		for {
			signal := <-slugUpdateChannel
			signal.Err <- resolveUpdateTitleSlugSignal(signal.Db, signal.SlugGroup)
		}
	}()
	return nil
}

func resolveUpdateTitleSlugSignal(db dbx.Builder, slugGroup string) error {
	titleTableName := (&models.Title{}).TableName()
	duplicatedTitles := []*models.Title{}
	if err := models.TitleQuery(db).
		Where(
			dbx.HashExp{
				fmt.Sprintf("%s.slugGroup", titleTableName): slugGroup,
			},
		).
		All(&duplicatedTitles); err != nil {
		return err
	}
	if len(duplicatedTitles) <= 0 {
		return nil
	}
	if len(duplicatedTitles) <= 1 {
		duplicatedTitles[0].Slug = duplicatedTitles[0].SlugGroup
		if err := db.Model(duplicatedTitles[0]).Update(); err != nil {
			return err
		}
		return nil
	}
	formatMap := map[string][]*models.Title{}
	for _, title := range duplicatedTitles {
		if err := title.Expand(db, models.ExpandMap{"format": {}}); err != nil {
			return err
		}
		if title.Format == nil {
			return errors.New("Expanded title's format should not be nil.")
		}
		formatMap[title.Format.Slug] = append(
			formatMap[title.Format.Slug],
			title,
		)
	}
	for format, titles := range formatMap {
		if len(titles) <= 1 {
			titles[0].Slug = fmt.Sprintf(
				"%s-%s",
				titles[0].SlugGroup,
				format,
			)
			if err := db.Model(titles[0]).Update(); err != nil {
				return err
			}
			continue
		}

		for _, title := range titles {
			singleFormatSlug := fmt.Sprintf(
				"%s-%s",
				title.SlugGroup,
				format,
			)
			switch title.Slug {
			case "", title.SlugGroup, singleFormatSlug:
				// do nothing

			default:
				continue
			}
			randomSlugId := strings.ToLower(tools.RandStr(5))
			title.Slug = fmt.Sprintf("%s-%s", title.SlugGroup, randomSlugId)
			if err := db.Model(title).Update(); err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateTitleSlug(db dbx.Builder, slugGroup string) error {
	err := make(chan error)
	slugUpdateChannel <- slugUpdateSignal{
		Db:        db,
		SlugGroup: slugGroup,
		Err:       err,
	}
	return <-err
}

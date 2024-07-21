package migrations

import (
	"regexp"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		releases := []*models.Release{}
		if err := models.ReleaseQuery(dao).All(&releases); err != nil {
			return err
		}
		for _, release := range releases {
			title := &models.Title{}
			if err := models.TitleQuery(dao).Where(&dbx.HashExp{
				"id": release.TitleId,
			}).One(title); err != nil {
				return err
			}

			// default title.name into release.name
			release.Name = title.Name

			if strings.Contains(release.Type, "Bản in đầu") {
				release.Type = ""
			}

			// turn anything inside bracket into disambiguation
			if strings.Contains(release.Type, "(") {
				reg := regexp.MustCompile(`(?s)\((.*)\)`)
				match := reg.FindAllStringSubmatch(release.Type, -1)
				release.Disambiguation = match[0][1]
				release.Type = reg.ReplaceAllString(release.Type, "")
			}

			if err := dao.Save(release); err != nil {
				return err
			}
		}

		return nil
	}, func(db dbx.Builder) error {
		return nil
	})
}

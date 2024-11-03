package migrations

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func findTitleDefaultRelease(dao *daos.Dao, title *models.Title) (*models.Release, error) {
	release := &models.Release{}
	err := models.ReleaseQuery(dao).Where(&dbx.HashExp{
		"title": title.Id,
	}).One(release)
	if err == nil {
		return release, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return nil, nil
}

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		titles := []*models.Title{}
		if err := models.TitleQuery(dao).All(&titles); err != nil {
			return err
		}
		for _, title := range titles {
			defaultRelease, err := findTitleDefaultRelease(dao, title)
			if err != nil {
				return err
			}
			if defaultRelease == nil {
				continue
			}
			title.DefaultReleaseId = defaultRelease.Id
			if err := dao.Save(title); err != nil {
				return err
			}
		}
		return nil
	}, func(db dbx.Builder) error {
		return nil
	})
}

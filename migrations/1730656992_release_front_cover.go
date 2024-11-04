package migrations

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func findReleaseFrontCover(dao *daos.Dao, release *models.Release) (*models.Asset, error) {
	asset := &models.Asset{}
	err := models.AssetQuery(dao).
		InnerJoin("books", dbx.NewExp("books.id = assets.book")).
		InnerJoin("publications", dbx.NewExp("publications.id = books.publication")).
		InnerJoin("releases", dbx.NewExp("releases.id = publications.release")).
		Where(&dbx.HashExp{
			"releases.id": release.Id,
		}).
		OrderBy("publications.volume ASC", "books.edition ASC", "priority ASC").
		One(asset)
	if err == nil {
		return asset, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return nil, nil
}

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		releases := []*models.Release{}
		if err := models.ReleaseQuery(dao).All(&releases); err != nil {
			return err
		}
		for _, release := range releases {
			defaultFront, err := findReleaseFrontCover(dao, release)
			if err != nil {
				return err
			}
			if defaultFront == nil {
				continue
			}
			release.FrontId = defaultFront.Id
			if err := dao.Save(release); err != nil {
				return err
			}
		}
		return nil
	}, func(db dbx.Builder) error {
		return nil
	})
}

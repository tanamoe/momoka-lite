package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func createCoverAsset(
	db dbx.Builder,
	bookId string,
	image string,
	priority int,
) error {
	coverAsset := &models.Asset{
		Id:       core.GenerateDefaultRandomId(),
		BookId:   bookId,
		TypeId:   models.AssetTypeCoverID,
		Image:    image,
		Priority: priority,
	}
	return db.Model(coverAsset).Insert()
}

func init() {
	m.Register(func(app core.App) error {
		resizedState := &models.State{
			Value: "0",
		}
		resizedState.Id = models.AssetImageResizedStateId
		state, err := models.FindStateById(app.DB(), models.AssetImageResizedStateId)
		if err != nil {
			return err
		}
		if state == nil {
			if err := app.DB().Model(resizedState).Insert(); err != nil {
				return err
			}
		} else {
			if err := app.DB().Model(resizedState).Update(); err != nil {
				return err
			}
		}
		books := []*models.Book{}
		if err := models.BookQuery(app.DB()).Select("id", "covers").All(&books); err != nil {
			return err
		}
		for _, book := range books {
			for id, cover := range book.Covers {
				if err := createCoverAsset(app.DB(), book.Id, cover, id+1); err != nil {
					return err
				}
			}
		}
		return nil
	}, func(app core.App) error {
		return nil
	})
}

package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func createCoverAsset(
	dao *daos.Dao,
	bookId string,
	image string,
	priority int,
) error {
	coverAsset := &models.Asset{
		BookId:   bookId,
		TypeId:   models.AssetTypeCoverID,
		Image:    image,
		Priority: priority,
	}
	return dao.Save(coverAsset)
}

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		resizedState := &models.State{
			Value: "0",
		}
		resizedState.Id = models.AssetImageResizedStateId
		if err := dao.Save(resizedState); err != nil {
			return err
		}
		books := []*models.Book{}
		if err := models.BookQuery(dao).Select("id", "covers").All(&books); err != nil {
			return err
		}
		for _, book := range books {
			for id, cover := range book.Covers {
				if err := createCoverAsset(dao, book.Id, cover, id+1); err != nil {
					return err
				}
			}
		}
		return nil
	}, func(db dbx.Builder) error {
		return nil
	})
}

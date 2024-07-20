package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		coverAssetType := &models.AssetType{
			Name: "Cover",
		}
		coverAssetType.Id = models.AssetTypeCoverID
		return dao.Save(coverAssetType)
	}, func(db dbx.Builder) error {
		return nil
	})
}

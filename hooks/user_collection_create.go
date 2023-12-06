package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
)

func registerOnUserCollectionCreate(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.
		OnModelAfterCreate((&models.Collection{}).TableName()).
		Add(
			func(e *core.ModelEvent) error {
				collectionId := e.Model.GetId()
				collection, err := models.FindCollectionById(app.Dao(), collectionId)
				if err != nil {
					return err
				}
				if err := app.Dao().Save(&models.CollectionMember{
					CollectionId: collection.Id,
					UserId:       collection.OwnerId,
					Role:         models.CollectionEditorRole,
				}); err != nil {
					return err
				}
				return nil
			},
		)
	return nil
}

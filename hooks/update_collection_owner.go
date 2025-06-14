package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
)

func registerOnUpdateCollectionOwner(
	app *pocketbase.PocketBase,
) error {
	app.
		OnModelAfterCreateSuccess((&models.Collection{}).TableName()).
		BindFunc(
			func(e *core.ModelEvent) error {
				collectionId := e.Model.PK().(string)
				collection, err := models.FindCollectionById(app.DB(), collectionId)

				if err != nil {
					return err
				}
				if err := collection.AddMember(
					app.DB(),
					collection.OwnerId,
					models.CollectionEditorRole,
				); err != nil {
					return err
				}
				return e.Next()
			},
		)

	app.
		OnModelAfterUpdateSuccess((&models.Collection{}).TableName()).
		BindFunc(
			func(e *core.ModelEvent) error {
				collectionId := e.Model.PK().(string)
				collection, err := models.FindCollectionById(app.DB(), collectionId)
				if err != nil {
					return err
				}
				if err := collection.AddMember(
					app.DB(),
					collection.OwnerId,
					models.CollectionEditorRole,
				); err != nil {
					return err
				}
				return e.Next()
			},
		)
	return nil
}

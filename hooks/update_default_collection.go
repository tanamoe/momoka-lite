package hooks

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
)

func registerOnUpdateDefaultCollection(
	app *pocketbase.PocketBase,
) error {
	app.
		OnModelAfterCreateSuccess((&models.Collection{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			if err := onUpdateDefaultCollection(app, e); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterUpdateSuccess((&models.Collection{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			if err := onUpdateDefaultCollection(app, e); err != nil {
				return err
			}
			return e.Next()
		})
	return nil
}

func onUpdateDefaultCollection(
	app *pocketbase.PocketBase,
	e *core.ModelEvent,
) error {
	collectionId := e.Model.PK().(string)
	collection, err := models.FindCollectionById(app.DB(), collectionId)
	if err != nil {
		return err
	}
	if !collection.Default {
		return nil
	}
	if err := markOtherOwnedCollectionAsNotDefault(
		app,
		collection,
	); err != nil {
		return err
	}
	return nil
}

func markOtherOwnedCollectionAsNotDefault(
	app *pocketbase.PocketBase,
	collection *models.Collection,
) error {
	return app.RunInTransaction(
		func(app core.App) error {
			collections := []*models.Collection{}
			if err := models.CollectionQuery(app.DB()).
				Where(
					dbx.And(
						dbx.HashExp{"owner": collection.OwnerId},
						dbx.Not(dbx.HashExp{"id": collection.Id}),
					),
				).All(&collections); err != nil {
				return err
			}

			for _, collection := range collections {
				collection.Default = false
				if err := app.UnsafeWithoutHooks().DB().Model(collection).Update(); err != nil {
					return nil
				}
			}
			return nil
		},
	)
}

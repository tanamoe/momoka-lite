package hooks

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"tana.moe/momoka-lite/models"
)

func registerOnUpdateDefaultCollection(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.
		OnModelAfterCreate((&models.Collection{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			return onUpdateDefaultCollection(app, context, e)
		})

	app.
		OnModelAfterUpdate((&models.Collection{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			return onUpdateDefaultCollection(app, context, e)
		})
	return nil
}

func onUpdateDefaultCollection(
	app *pocketbase.PocketBase,
	context *models.AppContext,
	e *core.ModelEvent,
) error {
	collectionId := e.Model.GetId()
	collection, err := models.FindCollectionById(app.Dao(), collectionId)
	if err != nil {
		return err
	}
	if !collection.Default {
		return nil
	}
	if err := markOtherOwnedCollectionAsNotDefault(
		app.Dao(),
		collection,
	); err != nil {
		return err
	}
	return nil
}

func markOtherOwnedCollectionAsNotDefault(
	dao *daos.Dao,
	collection *models.Collection,
) error {
	return dao.RunInTransaction(
		func(dao *daos.Dao) error {
			collections := []*models.Collection{}
			if err := models.CollectionQuery(dao).
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
				if err := dao.WithoutHooks().Save(collection); err != nil {
					return nil
				}
			}
			return nil
		},
	)
}

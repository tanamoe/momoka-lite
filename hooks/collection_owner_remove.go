package hooks

import (
	"errors"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
)

var collectionOwnerRemoveError = errors.New("Removing owner from a collection is forbidden.")

func registerOnCollectionOwnerRemove(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.
		OnModelAfterDelete((&models.CollectionMember{}).TableName()).
		Add(
			func(e *core.ModelEvent) error {
				collectionMember := e.Model.(*models.CollectionMember)
				if err := collectionMember.Expand(app.Dao(), models.ExpandMap{
					"collection": {},
				}); err != nil {
					return err
				}
				if collectionMember.Collection == nil {
					return nil
				}
				if collectionMember.UserId == collectionMember.Collection.OwnerId {
					return collectionOwnerRemoveError
				}
				return nil
			},
		)
	return nil
}

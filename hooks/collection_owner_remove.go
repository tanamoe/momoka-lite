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
) error {
	app.
		OnModelDelete((&models.CollectionMember{}).TableName()).
		BindFunc(
			func(e *core.ModelEvent) error {
				collectionMemberId := e.Model.PK().(string)
				collectionMember, err := models.FindCollectionMemberById(app.DB(), collectionMemberId)
				if err != nil {
					return err
				}
				if err := collectionMember.Expand(app.DB(), models.ExpandMap{
					"collection": {},
				}); err != nil {
					return err
				}
				if collectionMember.Collection == nil {
					return e.Next()
				}
				if collectionMember.UserId == collectionMember.Collection.OwnerId {
					return collectionOwnerRemoveError
				}
				return e.Next()
			},
		)
	return nil
}

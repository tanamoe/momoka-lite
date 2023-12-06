package hooks

import (
	"github.com/pocketbase/pocketbase"
	"tana.moe/momoka-lite/models"
)

func RegisterHooks(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	if err := registerAppendImageSecretHook(app, context); err != nil {
		return err
	}

	if err := registerOnUserCollectionCreate(app, context); err != nil {
		return err
	}

	if err := registerOnCollectionOwnerRemove(app, context); err != nil {
		return err
	}

	return nil
}

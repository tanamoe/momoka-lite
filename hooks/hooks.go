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

	if err := registerOnUpdateCollectionOwner(app, context); err != nil {
		return err
	}

	if err := registerOnCollectionOwnerRemove(app, context); err != nil {
		return err
	}

	if err := registerOnUpdateDefaultCollection(app, context); err != nil {
		return err
	}

	if err := registerAppendInCollectionsMetadataHook(app, context); err != nil {
		return err
	}

	if err := registerUpdateTitleSlugHook(app, context); err != nil {
		return err
	}

	if err := registerOnTitleIndexShouldChangeHook(app, context); err != nil {
		return err
	}

	return nil
}

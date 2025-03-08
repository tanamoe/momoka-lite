package hooks

import (
	"github.com/pocketbase/pocketbase"
)

func RegisterHooks(
	app *pocketbase.PocketBase,
) error {
	if err := registerAppendImageSecretHook(app); err != nil {
		return err
	}

	if err := registerOnUpdateCollectionOwner(app); err != nil {
		return err
	}

	if err := registerOnCollectionOwnerRemove(app); err != nil {
		return err
	}

	if err := registerOnUpdateDefaultCollection(app); err != nil {
		return err
	}

	if err := registerAppendInCollectionsMetadataHook(app); err != nil {
		return err
	}

	if err := registerUpdateTitleSlugHook(app); err != nil {
		return err
	}

	if err := registerOnTitleIndexShouldChangeHook(app); err != nil {
		return err
	}

	if err := registerOnReleaseIndexShouldChangeHook(app); err != nil {
		return err
	}

	if err := registerResizeAssetImageHook(app); err != nil {
		return err
	}

	if err := registerPublicationUpdateDefaultBook(app); err != nil {
		return err
	}

	return nil
}

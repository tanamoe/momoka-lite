package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services/chie"
)

func registerOnReleaseIndexShouldChangeHook(
	app *pocketbase.PocketBase,
) error {
	app.
		OnModelAfterCreateSuccess((&models.Release{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			release := &models.Release{}
			release.Id = e.Model.PK().(string)
			if err := updateReleaseIndex(app, release); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterUpdateSuccess((&models.Release{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			release := &models.Release{}
			release.Id = e.Model.PK().(string)
			if err := updateReleaseIndex(app, release); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterDeleteSuccess((&models.Release{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			release := &models.Release{}
			release.Id = e.Model.PK().(string)
			if err := updateReleaseIndex(app, release); err != nil {
				return err
			}
			return e.Next()
		})

	return nil
}

func updateReleaseIndex(
	app *pocketbase.PocketBase,
	release *models.Release,
) error {
	if err := chie.UpdateReleaseIndex(app.DB(), release); err != nil {
		return err
	}
	return nil
}

package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services/chie"
)

func registerOnReleaseIndexShouldChangeHook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.
		OnModelAfterCreate((&models.Release{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			release := &models.Release{}
			release.Id = e.Model.GetId()
			return updateReleaseIndex(app, context, release)
		})

	app.
		OnModelAfterUpdate((&models.Release{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			release := &models.Release{}
			release.Id = e.Model.GetId()
			return updateReleaseIndex(app, context, release)
		})

	app.
		OnModelAfterDelete((&models.Release{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			release := &models.Release{}
			release.Id = e.Model.GetId()
			return updateReleaseIndex(app, context, release)
		})

	return nil
}

func updateReleaseIndex(
	app *pocketbase.PocketBase,
	context *models.AppContext,
	release *models.Release,
) error {
	if err := chie.UpdateReleaseIndex(app.Dao(), release); err != nil {
		return err
	}
	return nil
}

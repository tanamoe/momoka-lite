package services

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services/chie"
)

func Start(app *pocketbase.PocketBase, context *models.AppContext) error {
	if err := startPreMigrationService(app, context); err != nil {
		return err
	}
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		if err := startPostMigrationService(app, context); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func startPreMigrationService(app *pocketbase.PocketBase, context *models.AppContext) error {
	if err := startUpdateSlugService(app, context); err != nil {
		return err
	}
	return nil
}

func startPostMigrationService(app *pocketbase.PocketBase, context *models.AppContext) error {
	if err := chie.StartService(app, context); err != nil {
		return err
	}
	return nil
}

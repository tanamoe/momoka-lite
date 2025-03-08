package services

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/services/chie"
)

func Start(app *pocketbase.PocketBase) error {
	if err := startPreMigrationService(app); err != nil {
		return err
	}
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		if err := startPostMigrationService(app); err != nil {
			return err
		}
		return e.Next()
	})
	return nil
}

func startPreMigrationService(app *pocketbase.PocketBase) error {
	if err := startUpdateSlugService(app); err != nil {
		return err
	}
	return nil
}

func startPostMigrationService(app *pocketbase.PocketBase) error {
	if err := chie.StartService(app); err != nil {
		return err
	}
	return nil
}

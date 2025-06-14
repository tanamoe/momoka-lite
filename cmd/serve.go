package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"tana.moe/momoka-lite/apis"
	"tana.moe/momoka-lite/hooks"
	_ "tana.moe/momoka-lite/migrations"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services"
	"tana.moe/momoka-lite/tools"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	app := pocketbase.New()

	if err := initStorage(app); err != nil {
		log.Fatal(err)
		return
	}

	if err := registerMiddleware(app); err != nil {
		log.Fatal(err)
		return
	}

	if err := registerApis(app); err != nil {
		log.Fatal(err)
		return
	}

	if err := hooks.RegisterHooks(app); err != nil {
		log.Fatal(err)
		return
	}

	if err := startServices(app); err != nil {
		log.Fatal(err)
		return
	}

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: false,
	})

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		if err := refreshAppState(app); err != nil {
			panic(err)
		}
		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
		return
	}
}

func initStorage(
	app *pocketbase.PocketBase,
) error {
	app.Store().Set(models.ImagorSecretKey, os.Getenv(models.ImagorSecretKey))
	return nil
}

func registerMiddleware(
	app *pocketbase.PocketBase,
) error {
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		return e.Next()
	})
	return nil
}

func registerApis(
	app *pocketbase.PocketBase,
) error {
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		apis.RegisterApis(app, e)
		return e.Next()
	})
	return nil
}

func startServices(
	app *pocketbase.PocketBase,
) error {
	if err := services.Start(app); err != nil {
		return err
	}
	return nil
}

func refreshAppState(
	app *pocketbase.PocketBase,
) error {
	db := app.DB()
	states := map[string]string{
		models.ImagorSecretStateId:      imagorSecretState(app),
		models.AssetImageResizedStateId: "1",
	}
	for id, stateValue := range states {
		state, err := models.FindStateById(db, id)
		if err != nil {
			return err
		}
		if state == nil {
			return errors.New(fmt.Sprintf("Unable to locate state `%s`", id))
		}
		if state.Value == stateValue {
			continue
		}
		state.Value = stateValue
		if err := db.Model(state).Update(); err != nil {
			return err
		}
	}
	return nil
}

func imagorSecretState(
	app *pocketbase.PocketBase,
) string {
	secret := app.Store().Get(models.ImagorSecretKey).(string)
	if secret == "" {
		return ""
	}
	return tools.SHA256(secret)
}

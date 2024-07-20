package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
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

	context, err := models.NewContext()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := registerMiddleware(app, context); err != nil {
		log.Fatal(err)
		return
	}

	if err := registerApis(app, context); err != nil {
		log.Fatal(err)
		return
	}

	if err := hooks.RegisterHooks(app, context); err != nil {
		log.Fatal(err)
		return
	}

	if err := startServices(app, context); err != nil {
		log.Fatal(err)
		return
	}

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: false,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		if err := refreshAppState(app, context); err != nil {
			panic(err)
		}
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
		return
	}
}

func registerMiddleware(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(appendAppContext(context))
		return nil
	})
	return nil
}

func registerApis(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		apis.RegisterApis(app, e)
		return nil
	})
	return nil
}

func appendAppContext(context *models.AppContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(&models.EchoContext{
				Context:    c,
				AppContext: context,
			})
		}
	}
}

func startServices(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	if err := services.Start(app, context); err != nil {
		return err
	}
	return nil
}

func refreshAppState(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	dao := app.Dao()
	states := map[string]string{
		models.ImagorSecretStateId:      imagorSecretState(context),
		models.AssetImageResizedStateId: "1",
	}
	for id, stateValue := range states {
		state, err := models.FindStateById(dao, id)
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
		if err := dao.Save(state); err != nil {
			return err
		}
	}
	return nil
}

func imagorSecretState(
	context *models.AppContext,
) string {
	if context.ImagorSecret == "" {
		return ""
	}
	return tools.SHA256(context.ImagorSecret)
}

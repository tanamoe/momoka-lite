package main

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"tana.moe/momoka-lite/apis"
	"tana.moe/momoka-lite/hooks"
	_ "tana.moe/momoka-lite/migrations"
	"tana.moe/momoka-lite/models"
)

func main() {
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

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: false,
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

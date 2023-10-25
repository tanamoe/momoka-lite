package hooks

import (
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
)

func registerDocsHandler(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./docs/.vitepress/dist"), false))
		return nil
	})
	return nil
}

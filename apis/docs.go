package apis

import (
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func registerDocsRoute(app *pocketbase.PocketBase, e *core.ServeEvent) error {
	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./docs/.vitepress/dist"), false))
	return nil
}

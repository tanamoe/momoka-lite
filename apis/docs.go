package apis

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"tana.moe/momoka-lite/docs"
)

func registerDocsRoute(app *pocketbase.PocketBase, e *core.ServeEvent) error {
	e.Router.GET(
		"/*",
		apis.StaticDirectoryHandler(docs.DistDirFS, false),
		apis.ActivityLogger(app),
	)
	return nil
}

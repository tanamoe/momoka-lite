package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
)

func appendImageSecret(
	app *pocketbase.PocketBase,
	context *models.AppContext,
	e *core.RecordsListEvent,
) error {
	for _, record := range e.Records {
		cover := record.GetString("cover")

		record.Set("metadata", map[string]interface{}{
			"srcset": cover + "1280w",
		})
	}
	return nil
}

func registerAppendImageSecretHook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.OnRecordsListRequest().Add(func(e *core.RecordsListEvent) error {
		return appendImageSecret(app, context, e)
	})
	return nil
}

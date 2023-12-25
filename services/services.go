package services

import (
	"github.com/pocketbase/pocketbase"
	"tana.moe/momoka-lite/models"
)

func Start(app *pocketbase.PocketBase, context *models.AppContext) error {
	if err := startUpdateSlugService(app, context); err != nil {
		return err
	}
	return nil
}

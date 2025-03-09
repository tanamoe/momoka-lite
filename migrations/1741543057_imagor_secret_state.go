package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func init() {
	m.Register(func(app core.App) error {
		imagorSecretState := &models.State{
			Value: "",
		}
		imagorSecretState.Id = models.ImagorSecretStateId
		state, err := models.FindStateById(app.DB(), models.ImagorSecretStateId)
		if err != nil {
			return err
		}
		if state != nil {
			return app.DB().Model(imagorSecretState).Update()
		}
		return app.DB().Model(imagorSecretState).Insert()
	}, func(app core.App) error {
		return nil
	})
}

package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	pmodels "github.com/pocketbase/pocketbase/models"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services/chie"
)

func registerOnTitleIndexShouldChangeHook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.
		OnModelAfterCreate((&models.Title{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			title := &models.Title{}
			title.Id = e.Model.GetId()
			return updateTitleIndex(app, context, title)
		})

	app.
		OnModelAfterUpdate((&models.Title{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			title := &models.Title{}
			title.Id = e.Model.GetId()
			return updateTitleIndex(app, context, title)
		})

	app.
		OnModelAfterDelete((&models.Title{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			title := &models.Title{}
			title.Id = e.Model.GetId()
			return updateTitleIndex(app, context, title)
		})

	app.
		OnModelAfterCreate((&models.Work{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			work := e.Model.(*pmodels.Record)
			title := &models.Title{}
			title.Id = work.GetString("title")
			return updateTitleIndex(app, context, title)
		})

	app.
		OnModelAfterUpdate((&models.Work{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			work := e.Model.(*pmodels.Record)
			title := &models.Title{}
			title.Id = work.GetString("title")
			return updateTitleIndex(app, context, title)
		})

	app.
		OnModelAfterDelete((&models.Work{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			work := e.Model.(*pmodels.Record)
			title := &models.Title{}
			title.Id = work.GetString("title")
			return updateTitleIndex(app, context, title)
		})
	app.
		OnModelAfterCreate((&models.AdditionalTitleName{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			additionalName := e.Model.(*pmodels.Record)
			title := &models.Title{}
			title.Id = additionalName.GetString("title")
			return updateTitleIndex(app, context, title)
		})

	app.
		OnModelAfterUpdate((&models.AdditionalTitleName{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			additionalName := e.Model.(*pmodels.Record)
			title := &models.Title{}
			title.Id = additionalName.GetString("title")
			return updateTitleIndex(app, context, title)
		})

	app.
		OnModelAfterDelete((&models.AdditionalTitleName{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			additionalName := e.Model.(*pmodels.Record)
			title := &models.Title{}
			title.Id = additionalName.GetString("title")
			return updateTitleIndex(app, context, title)
		})

	return nil
}

func updateTitleIndex(
	app *pocketbase.PocketBase,
	context *models.AppContext,
	title *models.Title,
) error {
	if err := chie.UpdateTitleIndex(app.Dao(), title); err != nil {
		return err
	}
	return nil
}

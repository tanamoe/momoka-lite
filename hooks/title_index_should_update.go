package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services/chie"
)

func registerOnTitleIndexShouldChangeHook(
	app *pocketbase.PocketBase,
) error {
	app.
		OnModelAfterCreateSuccess((&models.Title{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			title := &models.Title{}
			title.Id = e.Model.PK().(string)
			if err := updateTitleIndex(app, title); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterUpdateSuccess((&models.Title{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			title := &models.Title{}
			title.Id = e.Model.PK().(string)
			if err := updateTitleIndex(app, title); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterDeleteSuccess((&models.Title{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			title := &models.Title{}
			title.Id = e.Model.PK().(string)
			if err := updateTitleIndex(app, title); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterCreateSuccess((&models.Work{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			work := e.Model.(*core.Record)
			title := &models.Title{}
			title.Id = work.GetString("title")
			if err := updateTitleIndex(app, title); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterUpdateSuccess((&models.Work{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			work := e.Model.(*core.Record)
			title := &models.Title{}
			title.Id = work.GetString("title")
			if err := updateTitleIndex(app, title); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterDeleteSuccess((&models.Work{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			work := e.Model.(*core.Record)
			title := &models.Title{}
			title.Id = work.GetString("title")
			if err := updateTitleIndex(app, title); err != nil {
				return err
			}
			return e.Next()
		})
	app.
		OnModelAfterCreateSuccess((&models.AdditionalTitleName{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			additionalName := e.Model.(*core.Record)
			title := &models.Title{}
			title.Id = additionalName.GetString("title")
			if err := updateTitleIndex(app, title); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterUpdateSuccess((&models.AdditionalTitleName{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			additionalName := e.Model.(*core.Record)
			title := &models.Title{}
			title.Id = additionalName.GetString("title")
			if err := updateTitleIndex(app, title); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterDeleteSuccess((&models.AdditionalTitleName{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			additionalName := e.Model.(*core.Record)
			title := &models.Title{}
			title.Id = additionalName.GetString("title")
			if err := updateTitleIndex(app, title); err != nil {
				return err
			}
			return e.Next()
		})

	return nil
}

func updateTitleIndex(
	app *pocketbase.PocketBase,
	title *models.Title,
) error {
	if err := chie.UpdateTitleIndex(app.DB(), title); err != nil {
		return err
	}
	return nil
}

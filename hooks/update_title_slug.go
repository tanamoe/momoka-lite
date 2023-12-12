package hooks

import (
	"github.com/gosimple/slug"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services"
)

func registerUpdateTitleSlugHook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.
		OnModelAfterCreate((&models.Title{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			return updateTitleSlug(app, context, e)
		})

	app.
		OnModelAfterUpdate((&models.Title{}).TableName()).
		Add(func(e *core.ModelEvent) error {
			return updateTitleSlug(app, context, e)
		})
	return nil
}

func updateTitleSlug(
	app *pocketbase.PocketBase,
	context *models.AppContext,
	e *core.ModelEvent,
) error {
	titleId := e.Model.GetId()
	title, err := models.FindTitleById(e.Dao, titleId)
	if err != nil {
		return err
	}
	if title.SlugGroup == "" {
		title.SlugGroup = slug.Make(title.Name)
		title.Slug = title.SlugGroup
		if err := e.Dao.WithoutHooks().Save(title); err != nil {
			return err
		}
	}
	if err := services.UpdateTitleSlug(
		e.Dao.WithoutHooks(),
		title.SlugGroup,
	); err != nil {
		return err
	}
	return nil
}

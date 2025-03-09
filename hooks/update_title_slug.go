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
) error {
	app.
		OnModelAfterCreateSuccess((&models.Title{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			if err := updateTitleSlug(app, e); err != nil {
				return err
			}
			return e.Next()
		})

	app.
		OnModelAfterUpdateSuccess((&models.Title{}).TableName()).
		BindFunc(func(e *core.ModelEvent) error {
			if err := updateTitleSlug(app, e); err != nil {
				return err
			}
			return e.Next()
		})
	return nil
}

func updateTitleSlug(
	app *pocketbase.PocketBase,
	e *core.ModelEvent,
) error {
	titleId := e.Model.PK().(string)
	title, err := models.FindTitleById(app.DB(), titleId)
	if err != nil {
		return err
	}
	if title.SlugGroup == "" {
		title.SlugGroup = slug.Make(title.Name)
		title.Slug = title.SlugGroup
		if err := app.UnsafeWithoutHooks().DB().Model(title).Update(); err != nil {
			return err
		}
	}
	if err := services.UpdateTitleSlug(
		app.UnsafeWithoutHooks().DB(),
		title.SlugGroup,
	); err != nil {
		return err
	}
	return nil
}

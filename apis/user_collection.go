package apis

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	pmodels "github.com/pocketbase/pocketbase/models"
	"tana.moe/momoka-lite/models"
)

func registerUserCollectionRoute(
	app *pocketbase.PocketBase,
	core *core.ServeEvent,
) error {
	core.Router.GET(
		"/api/user-collection/:collectionId",
		viewRouteHandler(app, core, onRequestUserCollectionById),
		apis.ActivityLogger(app),
	)
	return nil
}

func onRequestUserCollectionById(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	expand models.ExpandMap,
) (item *models.Collection, err error) {
	admin, _ := c.Get(apis.ContextAdminKey).(*pmodels.Admin)
	record, _ := c.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
	collectionId := c.PathParam("collectionId")
	collection, err := models.FindCollectionById(app.Dao(), collectionId)
	if err != nil {
		return nil, err
	}
	if collection == nil {
		return nil, notFoundError
	}
	if admin != nil {
		return collection, nil
	}
	canBeAccessed, err := collection.CanBeAccessedBy(app.Dao(), record.Id)
	if err != nil {
		return nil, err
	}
	if !canBeAccessed {
		return nil, notFoundError
	}
	return nil, notFoundError
}

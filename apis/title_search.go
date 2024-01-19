package apis

import (
	"database/sql"
	"errors"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services/chie"
)

func registerTitleSearchRoute(
	app *pocketbase.PocketBase,
	core *core.ServeEvent,
) error {
	core.Router.POST(
		"/api/collections/titles/browse",
		listRouteHandler(app, core, onTitleSearchRequest),
		apis.ActivityLogger(app),
	)
	return nil
}

func onTitleSearchRequest(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []*models.Title, rpage uint, rperPage int, totalItems uint, totalPages uint, err error) {
	if perPage <= 0 {
		perPage = 25
	}
	req := chie.TitleSearchRequest{}
	if err := c.Bind(&req); err != nil {
		return nil, page, perPage, 0, 0, errors.Join(invalidRequestError, err)
	}
	req.Offset = int(page-1) * perPage
	req.Limit = perPage
	res := chie.SearchForTitles(req)
	if res.Error != nil {
		return nil, page, perPage, 0, 0, res.Error
	}
	if res.TotalItems > 0 {
		totalPages = uint((res.TotalItems-1)/perPage) + 1
	}
	titlesMap := map[string]*models.Title{}
	titleIds := []any{}
	for _, titleId := range res.Items {
		titleIds = append(titleIds, titleId)
	}
	unsortedItems := []*models.Title{}
	err = models.TitleQuery(app.Dao()).Where(dbx.HashExp{"id": titleIds}).All(&unsortedItems)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, page, perPage, uint(res.TotalItems), totalPages, nil
	}
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	for _, item := range unsortedItems {
		titlesMap[item.Id] = item
	}
	for _, itemId := range res.Items {
		if item, exist := titlesMap[itemId]; exist {
			if err := item.Expand(app.Dao(), expand); err != nil {
				return nil, page, perPage, 0, 0, err
			}
			items = append(items, item)
		}
	}
	return items, page, perPage, uint(res.TotalItems), totalPages, nil
}

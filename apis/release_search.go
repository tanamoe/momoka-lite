package apis

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services/chie"
)

func registerReleaseSearchRoute(
	app *pocketbase.PocketBase,
	core *core.ServeEvent,
) error {
	core.Router.POST(
		"/api/collections/releases/browse",
		listRouteHandler(app, core, onReleaseSearchRequest),
	)
	return nil
}

func onReleaseSearchRequest(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []*models.Release, rpage uint, rperPage int, totalItems uint, totalPages uint, err error) {
	if perPage <= 0 {
		perPage = 25
	}
	req := chie.ReleaseSearchRequest{}
	if err := e.BindBody(&req); err != nil {
		return nil, page, perPage, 0, 0, errors.Join(invalidRequestError, err)
	}
	req.Offset = int(page-1) * perPage
	req.Limit = perPage
	res := chie.SearchForReleases(req)
	if res.Error != nil {
		return nil, page, perPage, 0, 0, res.Error
	}
	if res.TotalItems > 0 {
		totalPages = uint((res.TotalItems-1)/perPage) + 1
	}
	releasesMap := map[string]*models.Release{}
	releaseIds := []any{}
	for _, releaseId := range res.Items {
		releaseIds = append(releaseIds, releaseId)
	}
	unsortedItems := []*models.Release{}
	err = models.ReleaseQuery(app.DB()).Where(dbx.HashExp{"id": releaseIds}).All(&unsortedItems)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, page, perPage, uint(res.TotalItems), totalPages, nil
	}
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	for _, item := range unsortedItems {
		releasesMap[item.Id] = item
	}
	for _, itemId := range res.Items {
		if item, exist := releasesMap[itemId]; exist {
			if err := item.Expand(app.DB(), expand); err != nil {
				return nil, page, perPage, 0, 0, err
			}
			items = append(items, item)
		}
	}
	return items, page, perPage, uint(res.TotalItems), totalPages, nil
}

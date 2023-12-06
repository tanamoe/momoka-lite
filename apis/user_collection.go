package apis

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
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
	core.Router.GET(
		"/api/user-collection/:collectionId/books",
		listRouteHandler(app, core, onRequestBooksInCollection),
		apis.ActivityLogger(app),
	)
	core.Router.GET(
		"/api/user-collection/:collectionId/members",
		listRouteHandler(app, core, onRequestMembersInCollection),
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
	if collectionId == "default" {
		return onRequestUserDefaultCollection(app, e, c, expand)
	}
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
	userId := ""
	if record != nil {
		userId = record.Id
	}
	canBeAccessed, err := collection.CanBeAccessedBy(app.Dao(), userId)
	if err != nil {
		return nil, err
	}
	if !canBeAccessed {
		return nil, notFoundError
	}
	if err := collection.Expand(app.Dao(), expand); err != nil {
		return nil, err
	}
	return collection, nil
}

func onRequestUserDefaultCollection(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	expand models.ExpandMap,
) (item *models.Collection, err error) {
	record, _ := c.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
	if record == nil {
		return nil, unauthorizedError
	}
	collection, err := models.FindUserDefaultCollection(app.Dao(), record.Id)
	if err != nil {
		return nil, err
	}
	if collection == nil {
		return nil, notFoundError
	}
	if err := collection.Expand(app.Dao(), expand); err != nil {
		return nil, err
	}
	return collection, nil
}

func onRequestBooksInCollection(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []*models.CollectionBook, rpage uint, rperPage int, totalItems uint, totalPages uint, err error) {
	admin, _ := c.Get(apis.ContextAdminKey).(*pmodels.Admin)
	record, _ := c.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
	collectionId := c.PathParam("collectionId")
	collection, err := models.FindCollectionById(app.Dao(), collectionId)
	if perPage <= 0 {
		perPage = 25
	}
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	if collection == nil {
		return nil, page, perPage, 0, 0, err
	}
	if admin != nil {
		items, err = fetchBooksInCollection(app.Dao(), collection, page, perPage)
	} else {
		userId := ""
		if record != nil {
			userId = record.Id
		}
		canBeAccessBy, err := collection.CanBeAccessedBy(app.Dao(), userId)
		if err != nil {
			return nil, page, perPage, 0, 0, err
		}
		if !canBeAccessBy {
			return nil, page, perPage, 0, 0, notFoundError
		}
		items, err = fetchBooksInCollection(app.Dao(), collection, page, perPage)
	}
	if err != nil {
		return nil, page, perPage, 0, 0, notFoundError
	}
	totalItems, err = countBooksInCollection(app.Dao(), collection)
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	for _, item := range items {
		if err := item.Expand(app.Dao(), expand); err != nil {
			return nil, page, perPage, 0, 0, err
		}
	}
	totalPages = uint((int(totalItems) + perPage - 1) / perPage)
	return items, page, perPage, totalItems, totalPages, nil
}

func onRequestMembersInCollection(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []*models.CollectionMember, rpage uint, rperPage int, totalItems uint, totalPages uint, err error) {
	admin, _ := c.Get(apis.ContextAdminKey).(*pmodels.Admin)
	record, _ := c.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
	collectionId := c.PathParam("collectionId")
	collection, err := models.FindCollectionById(app.Dao(), collectionId)
	if perPage <= 0 {
		perPage = 25
	}
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	if collection == nil {
		return nil, page, perPage, 0, 0, err
	}
	if admin != nil {
		items, err = fetchMembersInCollection(app.Dao(), collection, page, perPage)
	} else {
		userId := ""
		if record != nil {
			userId = record.Id
		}
		canBeAccessBy, err := collection.CanBeAccessedBy(app.Dao(), userId)
		if err != nil {
			return nil, page, perPage, 0, 0, err
		}
		if !canBeAccessBy {
			return nil, page, perPage, 0, 0, notFoundError
		}
		items, err = fetchMembersInCollection(app.Dao(), collection, page, perPage)
	}
	if err != nil {
		return nil, page, perPage, 0, 0, notFoundError
	}
	totalItems, err = countMembersInCollection(app.Dao(), collection)
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	for _, item := range items {
		if err := item.Expand(app.Dao(), expand); err != nil {
			return nil, page, perPage, 0, 0, err
		}
	}
	totalPages = uint((int(totalItems) + perPage - 1) / perPage)
	return items, page, perPage, totalItems, totalPages, nil
}

func booksInCollectionQuery(
	dao *daos.Dao,
	collection *models.Collection,
) *dbx.SelectQuery {
	return models.CollectionBookQuery(dao).
		AndWhere(dbx.HashExp{
			"collection": collection.Id,
		})
}

func fetchBooksInCollection(
	dao *daos.Dao,
	collection *models.Collection,
	page uint,
	perPage int,
) (items []*models.CollectionBook, err error) {
	items = []*models.CollectionBook{}
	err = booksInCollectionQuery(dao, collection).
		Limit(int64(perPage)).
		Offset(int64(page-1) * int64(perPage)).
		All(&items)
	return
}

func countBooksInCollection(
	dao *daos.Dao,
	collection *models.Collection,
) (count uint, err error) {
	type countData struct {
		Count uint `db:"count"`
	}
	result := &countData{
		Count: 0,
	}
	err = booksInCollectionQuery(dao, collection).
		Select("COUNT(id) AS count").
		One(&result)
	return result.Count, err
}

func membersInCollectionQuery(
	dao *daos.Dao,
	collection *models.Collection,
) *dbx.SelectQuery {
	return models.CollectionMemberQuery(dao).
		AndWhere(dbx.HashExp{
			"collection": collection.Id,
		})
}

func fetchMembersInCollection(
	dao *daos.Dao,
	collection *models.Collection,
	page uint,
	perPage int,
) (items []*models.CollectionMember, err error) {
	items = []*models.CollectionMember{}
	err = membersInCollectionQuery(dao, collection).
		Limit(int64(perPage)).
		Offset(int64(page-1) * int64(perPage)).
		All(&items)
	return
}

func countMembersInCollection(
	dao *daos.Dao,
	collection *models.Collection,
) (count uint, err error) {
	type countData struct {
		Count uint `db:"count"`
	}
	result := &countData{
		Count: 0,
	}
	err = membersInCollectionQuery(dao, collection).
		Select("COUNT(id) AS count").
		One(&result)
	return result.Count, err
}

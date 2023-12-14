package apis

import (
	"database/sql"
	"errors"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/forms"
	pmodels "github.com/pocketbase/pocketbase/models"
	"tana.moe/momoka-lite/models"
)

func registerUserCollectionRoute(
	app *pocketbase.PocketBase,
	core *core.ServeEvent,
) error {
	core.Router.POST(
		"/api/user-collection",
		upsertRouteHandler(app, core, onCollectionUpsertRequest),
		apis.ActivityLogger(app),
	)
	core.Router.GET(
		"/api/user-collection/:collectionId",
		viewRouteHandler(app, core, onRequestUserCollectionById),
		apis.ActivityLogger(app),
	)
	core.Router.POST(
		"/api/user-collection/:collectionId",
		upsertRouteHandler(app, core, onCollectionUpsertRequest),
		apis.ActivityLogger(app),
	)
	core.Router.DELETE(
		"/api/user-collection/:collectionId",
		deleteRouteHandler(app, core, onDeleteCollectionRequest),
		apis.ActivityLogger(app),
	)
	core.Router.GET(
		"/api/user-collection/:collectionId/books",
		listRouteHandler(app, core, onRequestBooksInCollection),
		apis.ActivityLogger(app),
	)
	core.Router.POST(
		"/api/user-collection/:collectionId/books",
		upsertRouteHandler(app, core, onUpsertBookToCollectionRequest),
		apis.ActivityLogger(app),
	)
	core.Router.DELETE(
		"/api/user-collection/:collectionId/books/:bookId",
		deleteRouteHandler(app, core, onDeleteBookFromCollectionRequest),
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

func onCollectionUpsertRequest(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	expand models.ExpandMap,
) (item *models.Collection, err error) {
	item = &models.Collection{}
	if err = c.Bind(item); err != nil {
		return nil, errors.Join(invalidRequestError, err)
	}
	admin, _ := c.Get(apis.ContextAdminKey).(*pmodels.Admin)
	record, _ := c.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
	if (admin == nil) && (record == nil) {
		return nil, unauthorizedError
	}
	item.Id = c.PathParam("collectionId")
	if (item.Id != "") && (admin == nil) {
		canEditCollection, err := item.CanBeEditedBy(app.Dao(), record.Id)
		if err != nil {
			return nil, err
		}
		canAccessCollection, err := item.CanBeAccessedBy(app.Dao(), record.Id)
		if err != nil {
			return nil, err
		}
		if !canEditCollection {
			if !canAccessCollection {
				return nil, notFoundError
			}
			return nil, forbiddenError
		}
	}
	if item.Id != "" {
		originalCollection, err := models.FindCollectionById(app.Dao(), item.Id)
		if err != nil {
			return nil, err
		}
		if originalCollection == nil {
			return nil, notFoundError
		}
		if item.OwnerId == "" {
			item.OwnerId = originalCollection.OwnerId
		}
		if item.OwnerId != originalCollection.OwnerId {
			if (admin == nil) && (record != nil) && (record.Id != originalCollection.OwnerId) {
				return nil, forbiddenError
			}
			newOwner, err := models.FindUserById(app.Dao(), item.OwnerId)
			if err != nil {
				return nil, err
			}
			if newOwner == nil {
				return nil, invalidRequestError
			}
			if admin == nil {
				item.Default = false
			}
		}
		if (admin == nil) && (record.Id != originalCollection.OwnerId) {
			if item.Default != originalCollection.Default {
				return nil, forbiddenError
			}
			if item.Order != originalCollection.Order {
				return nil, forbiddenError
			}
		}
	} else {
		if (admin != nil) && (item.OwnerId == "") {
			return nil, invalidRequestError
		}
		if record == nil {
			return nil, unauthorizedError
		}
		if record != nil {
			item.OwnerId = record.Id
		}
	}
	collection, err := app.Dao().FindCollectionByNameOrId((&models.Collection{}).TableName())
	if err != nil {
		return nil, err
	}
	r := pmodels.NewRecord(collection)
	if item.Id != "" {
		if r, err = app.Dao().FindRecordById((&models.Collection{}).TableName(), item.Id); err != nil {
			return nil, err
		}
	}
	form := forms.NewRecordUpsert(app, r)
	form.LoadData(map[string]any{
		"owner":       item.OwnerId,
		"visibility":  item.Visibility,
		"name":        item.Name,
		"default":     item.Default,
		"description": item.Description,
		"order":       item.Order,
	})
	if err := form.Submit(); err != nil {
		return nil, errors.Join(invalidRequestError, err)
	}
	if item, err = models.FindCollectionById(app.Dao(), r.Id); err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("Upserted collection is not suppose to be nil.")
	}
	if err = item.Expand(app.Dao(), expand); err != nil {
		return nil, err
	}
	return item, nil
}

func onDeleteCollectionRequest(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
) error {
	admin, _ := c.Get(apis.ContextAdminKey).(*pmodels.Admin)
	record, _ := c.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
	if (admin == nil) && (record == nil) {
		return unauthorizedError
	}
	collectionId := c.PathParam("collectionId")
	collection, err := models.FindCollectionById(app.Dao(), collectionId)
	if err != nil {
		return err
	}
	if collection == nil {
		return notFoundError
	}
	if admin == nil {
		if collection.OwnerId != record.Id {
			canAccessCollection, err := collection.CanBeAccessedBy(app.Dao(), record.Id)
			if err != nil {
				return err
			}
			if !canAccessCollection {
				return notFoundError
			}
			return forbiddenError
		}
	}
	return app.Dao().WithoutHooks().RunInTransaction(func(dao *daos.Dao) error {
		members := []*models.CollectionMember{}
		err := models.CollectionMemberQuery(dao).
			AndWhere(
				dbx.HashExp{
					"collection": collection.Id,
				},
			).
			All(&members)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		if err != nil {
			return err
		}
		for _, member := range members {
			if err := dao.Delete(member); err != nil {
				return err
			}
		}
		if err = dao.Delete(collection); err != nil {
			return err
		}
		return nil
	})
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

func onUpsertBookToCollectionRequest(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	expand models.ExpandMap,
) (item *models.CollectionBook, err error) {
	item = &models.CollectionBook{}
	if err = c.Bind(&item); err != nil {
		return nil, errors.Join(invalidRequestError, err)
	}
	admin, _ := c.Get(apis.ContextAdminKey).(*pmodels.Admin)
	record, _ := c.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
	if (admin == nil) && (record == nil) {
		return nil, unauthorizedError
	}
	item.CollectionId = c.PathParam("collectionId")
	if err := item.Expand(app.Dao(), models.ExpandMap{
		"collection": {},
		"book":       {},
	}); err != nil {
		return nil, err
	}
	if item.Collection == nil {
		return nil, notFoundError
	}
	if item.Book == nil {
		return nil, invalidRequestError
	}
	if admin == nil {
		canEditCollection, err := item.Collection.CanBeEditedBy(app.Dao(), record.Id)
		if err != nil {
			return nil, err
		}
		if !canEditCollection {
			canAccessCollection, err := item.Collection.CanBeAccessedBy(app.Dao(), record.Id)
			if err != nil {
				return nil, err
			}
			if !canAccessCollection {
				return nil, notFoundError
			}
			return nil, forbiddenError
		}
	}
	existItem, err := fetchBookInCollection(app.Dao(), item.Collection, item.Book)
	if err != nil {
		return nil, err
	}
	if existItem != nil {
		item.Id = existItem.Id
	}
	collection, err := app.Dao().FindCollectionByNameOrId(
		(&models.CollectionBook{}).TableName(),
	)
	if err != nil {
		return nil, err
	}
	r := pmodels.NewRecord(collection)
	if item.Id != "" {
		if r, err = app.Dao().FindRecordById((&models.CollectionBook{}).TableName(), item.Id); err != nil {
			return nil, err
		}
	}
	form := forms.NewRecordUpsert(app, r)
	form.LoadData(map[string]any{
		"collection": item.CollectionId,
		"book":       item.BookId,
		"quantity":   item.Quantity,
		"status":     item.Status,
		"notes":      item.Notes,
	})
	if err = form.Submit(); err != nil {
		return nil, errors.Join(invalidRequestError, err)
	}
	item, err = models.FindCollectionBookById(app.Dao(), r.Id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("Upserted book to collection not supposed to be nil.")
	}
	if err = item.Expand(app.Dao(), expand); err != nil {
		return nil, err
	}
	return item, nil
}

func onDeleteBookFromCollectionRequest(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
) error {
	item := &models.CollectionBook{}
	admin, _ := c.Get(apis.ContextAdminKey).(*pmodels.Admin)
	record, _ := c.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
	if (admin == nil) && (record == nil) {
		return unauthorizedError
	}
	item.CollectionId = c.PathParam("collectionId")
	item.BookId = c.PathParam("bookId")
	if err := item.Expand(app.Dao(), models.ExpandMap{
		"collection": {},
	}); err != nil {
		return err
	}
	if item.Collection == nil {
		return notFoundError
	}
	item.Book = &models.Book{}
	item.Book.Id = item.BookId
	if admin == nil {
		canEditCollection, err := item.Collection.CanBeEditedBy(app.Dao(), record.Id)
		if err != nil {
			return err
		}
		if !canEditCollection {
			canAccessCollection, err := item.Collection.CanBeAccessedBy(app.Dao(), record.Id)
			if err != nil {
				return err
			}
			if !canAccessCollection {
				return notFoundError
			}
			return forbiddenError
		}
	}
	existItem, err := fetchBookInCollection(app.Dao(), item.Collection, item.Book)
	if err != nil {
		return err
	}
	if existItem == nil {
		return nil
	}
	if err = app.Dao().Delete(existItem); err != nil {
		return err
	}
	return err
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

func fetchBookInCollection(
	dao *daos.Dao,
	collection *models.Collection,
	book *models.Book,
) (item *models.CollectionBook, err error) {
	item = &models.CollectionBook{}
	err = booksInCollectionQuery(dao, collection).
		AndWhere(
			dbx.HashExp{
				"collection": collection.Id,
				"book":       book.Id,
			},
		).
		One(&item)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return

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

package apis

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"tana.moe/momoka-lite/models"
)

func registerUserCollectionRoute(
	app *pocketbase.PocketBase,
	core *core.ServeEvent,
) error {
	core.Router.POST(
		"/api/user-collection",
		upsertRouteHandler(app, core, onCollectionUpsertRequest),
	)
	core.Router.GET(
		"/api/user-collection/{collectionId}",
		viewRouteHandler(app, core, onRequestUserCollectionById),
	)
	core.Router.POST(
		"/api/user-collection/{collectionId}",
		upsertRouteHandler(app, core, onCollectionUpsertRequest),
	)
	core.Router.DELETE(
		"/api/user-collection/{collectionId}",
		deleteRouteHandler(app, core, onDeleteCollectionRequest),
	)
	core.Router.GET(
		"/api/user-collection/{collectionId}/books",
		listRouteHandler(app, core, onRequestBooksInCollection),
	)
	core.Router.POST(
		"/api/user-collection/{collectionId}/books",
		upsertRouteHandler(app, core, onUpsertBookToCollectionRequest),
	)
	core.Router.DELETE(
		"/api/user-collection/{collectionId}/books/{bookId}",
		deleteRouteHandler(app, core, onDeleteBookFromCollectionRequest),
	)
	core.Router.GET(
		"/api/user-collection/{collectionId}/members",
		listRouteHandler(app, core, onRequestMembersInCollection),
	)
	return nil
}

func onRequestUserCollectionById(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	expand models.ExpandMap,
) (item *models.Collection, err error) {
	info, err := e.RequestInfo()
	if err != nil {
		return nil, err
	}
	collectionId := e.Request.PathValue("collectionId")
	if collectionId == "default" {
		return onRequestUserDefaultCollection(app, e, expand)
	}
	collection, err := models.FindCollectionById(app.DB(), collectionId)
	if err != nil {
		return nil, err
	}
	if collection == nil {
		return nil, notFoundError
	}
	if (info.Auth != nil) && (info.Auth.IsSuperuser()) {
		return collection, nil
	}
	userId := ""
	if info.Auth != nil {
		userId = info.Auth.Id
	}
	canBeAccessed, err := collection.CanBeAccessedBy(app.DB(), userId)
	if err != nil {
		return nil, err
	}
	if !canBeAccessed {
		return nil, notFoundError
	}
	if err := collection.Expand(app.DB(), expand); err != nil {
		return nil, err
	}
	return collection, nil
}

func onRequestUserDefaultCollection(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	expand models.ExpandMap,
) (item *models.Collection, err error) {
	info, err := e.RequestInfo()
	if err != nil {
		return nil, err
	}
	if info.Auth == nil {
		return nil, notFoundError
	}
	collection, err := models.FindUserDefaultCollection(app.DB(), info.Auth.Id)
	if err != nil {
		return nil, err
	}
	if collection == nil {
		return nil, notFoundError
	}
	if err := collection.Expand(app.DB(), expand); err != nil {
		return nil, err
	}
	return collection, nil
}

func onRequestBooksInCollection(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []*models.CollectionBook, rpage uint, rperPage int, totalItems uint, totalPages uint, err error) {
	info, err := e.RequestInfo()
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	collectionId := e.Request.PathValue("collectionId")
	collection, err := models.FindCollectionById(app.DB(), collectionId)
	if perPage <= 0 {
		perPage = 25
	}
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	if collection == nil {
		return nil, page, perPage, 0, 0, err
	}
	isAdmin := (info.Auth != nil) && (info.Auth.IsSuperuser())
	if isAdmin {
		items, err = fetchBooksInCollection(app.DB(), collection, page, perPage)
	} else {
		userId := ""
		if info.Auth != nil {
			userId = info.Auth.Id
		}
		canBeAccessBy, err := collection.CanBeAccessedBy(app.DB(), userId)
		if err != nil {
			return nil, page, perPage, 0, 0, err
		}
		if !canBeAccessBy {
			return nil, page, perPage, 0, 0, notFoundError
		}
		items, err = fetchBooksInCollection(app.DB(), collection, page, perPage)
	}
	if err != nil {
		return nil, page, perPage, 0, 0, notFoundError
	}
	totalItems, err = countBooksInCollection(app.DB(), collection)
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	for _, item := range items {
		if err := item.Expand(app.DB(), expand); err != nil {
			return nil, page, perPage, 0, 0, err
		}
	}
	totalPages = uint((int(totalItems) + perPage - 1) / perPage)
	return items, page, perPage, totalItems, totalPages, nil
}

func onCollectionUpsertRequest(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	expand models.ExpandMap,
) (item *models.Collection, err error) {
	item = &models.Collection{}
	if err = e.BindBody(item); err != nil {
		return nil, errors.Join(invalidRequestError, err)
	}
	info, err := e.RequestInfo()
	if err != nil {
		return nil, err
	}
	if info.Auth == nil {
		return nil, unauthorizedError
	}
	item.Id = e.Request.PathValue("collectionId")
	if (item.Id != "") && (!info.Auth.IsSuperuser()) {
		canEditCollection, err := item.CanBeEditedBy(app.DB(), info.Auth.Id)
		if err != nil {
			return nil, err
		}
		if !canEditCollection {
			if item.Visibility == models.CollectionPublic {
				return nil, forbiddenError
			}
			isMember, err := item.HadMember(app.DB(), info.Auth.Id)
			if err != nil {
				return nil, err
			}
			if !isMember {
				return nil, notFoundError
			}
			return nil, forbiddenError
		}
	}

	// Update collection
	if item.Id != "" {
		originalCollection, err := models.FindCollectionById(app.DB(), item.Id)
		if err != nil {
			return nil, err
		}
		if originalCollection == nil {
			return nil, notFoundError
		}
		if item.OwnerId == "" {
			item.OwnerId = originalCollection.OwnerId
		}

		// Transfering ownership of the collection
		// We do not expect non-admin user can assign it to be other's default collection
		if item.OwnerId != originalCollection.OwnerId {
			if (info.Auth != nil) && (info.Auth.Id != originalCollection.OwnerId) {
				return nil, forbiddenError
			}
			newOwner, err := models.FindUserById(app.DB(), item.OwnerId)
			if err != nil {
				return nil, err
			}
			if newOwner == nil {
				return nil, invalidRequestError
			}
			if !info.Auth.IsSuperuser() {
				item.Default = false
			}
		}

		// Either admin or the collection's owner can update collection order or default collection of user
		if (!info.Auth.IsSuperuser()) && (info.Auth.Id != originalCollection.OwnerId) {
			if item.Default != originalCollection.Default {
				return nil, forbiddenError
			}
			if item.Order != originalCollection.Order {
				return nil, forbiddenError
			}
		}
	} else { // Create collection
		if (!info.Auth.IsSuperuser()) && (item.OwnerId == "") {
			return nil, invalidRequestError
		}
		if !info.Auth.IsSuperuser() {
			item.OwnerId = info.Auth.Id
		}
	}
	collection, err := app.FindCollectionByNameOrId((&models.Collection{}).TableName())
	if err != nil {
		return nil, err
	}
	r := core.NewRecord(collection)
	if item.Id != "" {
		if r, err = app.FindRecordById((&models.Collection{}).TableName(), item.Id); err != nil {
			return nil, err
		}
	}
	form := forms.NewRecordUpsert(app, r)
	form.Load(map[string]any{
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
	if item, err = models.FindCollectionById(app.DB(), r.Id); err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("Upserted collection is not suppose to be nil.")
	}
	if err = item.Expand(app.DB(), expand); err != nil {
		return nil, err
	}
	return item, nil
}

func onDeleteCollectionRequest(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
) error {
	info, err := e.RequestInfo()
	if err != nil {
		return err
	}
	if info.Auth != nil {
		return unauthorizedError
	}
	collectionId := e.Request.PathValue("collectionId")
	collection, err := models.FindCollectionById(app.DB(), collectionId)
	if err != nil {
		return err
	}
	if collection == nil {
		return notFoundError
	}
	if !info.Auth.IsSuperuser() {
		if collection.OwnerId != info.Auth.Id {
			canAccessCollection, err := collection.CanBeAccessedBy(app.DB(), info.Auth.Id)
			if err != nil {
				return err
			}
			if !canAccessCollection {
				return notFoundError
			}
			return forbiddenError
		}
	}
	return app.UnsafeWithoutHooks().RunInTransaction(func(app core.App) error {
		members := []*models.CollectionMember{}
		err := models.CollectionMemberQuery(app.DB()).
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
			if err := app.DB().Model(member).Delete(); err != nil {
				return err
			}
		}
		if err = app.DB().Model(collection).Delete(); err != nil {
			return err
		}
		return nil
	})
}

func onRequestMembersInCollection(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []*models.CollectionMember, rpage uint, rperPage int, totalItems uint, totalPages uint, err error) {
	info, err := e.RequestInfo()
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	collectionId := e.Request.PathValue("collectionId")
	collection, err := models.FindCollectionById(app.DB(), collectionId)
	if perPage <= 0 {
		perPage = 25
	}
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	if collection == nil {
		return nil, page, perPage, 0, 0, err
	}
	if !info.Auth.IsSuperuser() {
		items, err = fetchMembersInCollection(app.DB(), collection, page, perPage)
	} else {
		userId := ""
		if info.Auth != nil {
			userId = info.Auth.Id
		}
		canBeAccessBy, err := collection.CanBeAccessedBy(app.DB(), userId)
		if err != nil {
			return nil, page, perPage, 0, 0, err
		}
		if !canBeAccessBy {
			return nil, page, perPage, 0, 0, notFoundError
		}
		items, err = fetchMembersInCollection(app.DB(), collection, page, perPage)
	}
	if err != nil {
		return nil, page, perPage, 0, 0, notFoundError
	}
	totalItems, err = countMembersInCollection(app.DB(), collection)
	if err != nil {
		return nil, page, perPage, 0, 0, err
	}
	for _, item := range items {
		if err := item.Expand(app.DB(), expand); err != nil {
			return nil, page, perPage, 0, 0, err
		}
	}
	totalPages = uint((int(totalItems) + perPage - 1) / perPage)
	return items, page, perPage, totalItems, totalPages, nil
}

func onUpsertBookToCollectionRequest(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	expand models.ExpandMap,
) (item *models.CollectionBook, err error) {
	item = &models.CollectionBook{}
	if err = e.BindBody(&item); err != nil {
		return nil, errors.Join(invalidRequestError, err)
	}
	info, err := e.RequestInfo()
	if err != nil {
		return nil, err
	}
	if info.Auth == nil {
		return nil, unauthorizedError
	}
	item.CollectionId = e.Request.PathValue("collectionId")
	if err := item.Expand(app.DB(), models.ExpandMap{
		"collection": nil,
		"book":       nil,
	}); err != nil {
		return nil, err
	}
	if item.Collection == nil {
		return nil, notFoundError
	}
	if item.Book == nil {
		return nil, invalidRequestError
	}
	if !info.Auth.IsSuperuser() {
		canEditCollection, err := item.Collection.CanBeEditedBy(app.DB(), info.Auth.Id)
		if err != nil {
			return nil, err
		}
		if !canEditCollection {
			if item.Collection.Visibility == models.CollectionPublic {
				return nil, forbiddenError
			}
			isMember, err := item.Collection.HadMember(app.DB(), info.Auth.Id)
			if err != nil {
				return nil, err
			}
			if !isMember {
				return nil, notFoundError
			}
			return nil, forbiddenError
		}
	}
	existItem, err := fetchBookInCollection(app.DB(), item.Collection, item.Book)
	if err != nil {
		return nil, err
	}
	if existItem != nil {
		item.Id = existItem.Id
	}
	collection, err := app.FindCollectionByNameOrId(
		(&models.CollectionBook{}).TableName(),
	)
	if err != nil {
		return nil, err
	}
	r := core.NewRecord(collection)
	if item.Id != "" {
		if r, err = app.FindRecordById((&models.CollectionBook{}).TableName(), item.Id); err != nil {
			return nil, err
		}
	}
	form := forms.NewRecordUpsert(app, r)
	form.Load(map[string]any{
		"collection": item.CollectionId,
		"book":       item.BookId,
		"quantity":   item.Quantity,
		"status":     item.Status,
		"notes":      item.Notes,
	})
	if err = form.Submit(); err != nil {
		return nil, errors.Join(invalidRequestError, err)
	}
	item, err = models.FindCollectionBookById(app.DB(), r.Id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("Upserted book to collection not supposed to be nil.")
	}
	if err = item.Expand(app.DB(), expand); err != nil {
		return nil, err
	}
	return item, nil
}

func onDeleteBookFromCollectionRequest(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
) error {
	item := &models.CollectionBook{}
	info, err := e.RequestInfo()
	if err != nil {
		return err
	}
	if info.Auth == nil {
		return unauthorizedError
	}
	item.CollectionId = e.Request.PathValue("collectionId")
	item.BookId = e.Request.PathValue("bookId")
	if err := item.Expand(app.DB(), models.ExpandMap{
		"collection": {},
	}); err != nil {
		return err
	}
	if item.Collection == nil {
		return notFoundError
	}
	item.Book = &models.Book{}
	item.Book.Id = item.BookId
	if !info.Auth.IsSuperuser() {
		canEditCollection, err := item.Collection.CanBeEditedBy(app.DB(), info.Auth.Id)
		if err != nil {
			return err
		}
		if !canEditCollection {
			if item.Collection.Visibility == models.CollectionPublic {
				return forbiddenError
			}
			isMember, err := item.Collection.HadMember(app.DB(), info.Auth.Id)
			if err != nil {
				return err
			}
			if !isMember {
				return notFoundError
			}
			return forbiddenError
		}
	}
	existItem, err := fetchBookInCollection(app.DB(), item.Collection, item.Book)
	if err != nil {
		return err
	}
	if existItem == nil {
		return nil
	}
	if err = app.DB().Model(existItem).Delete(); err != nil {
		return err
	}
	return err
}

func booksInCollectionQuery(
	db dbx.Builder,
	collection *models.Collection,
) *dbx.SelectQuery {
	return models.CollectionBookQuery(db).
		AndWhere(dbx.HashExp{
			"collection": collection.Id,
		})
}

func fetchBooksInCollection(
	db dbx.Builder,
	collection *models.Collection,
	page uint,
	perPage int,
) (items []*models.CollectionBook, err error) {
	items = []*models.CollectionBook{}
	err = booksInCollectionQuery(db, collection).
		Limit(int64(perPage)).
		Offset(int64(page-1) * int64(perPage)).
		All(&items)
	return
}

func countBooksInCollection(
	db dbx.Builder,
	collection *models.Collection,
) (count uint, err error) {
	type countData struct {
		Count uint `db:"count"`
	}
	result := &countData{
		Count: 0,
	}
	err = booksInCollectionQuery(db, collection).
		Select("COUNT(id) AS count").
		One(&result)
	return result.Count, err
}

func fetchBookInCollection(
	db dbx.Builder,
	collection *models.Collection,
	book *models.Book,
) (item *models.CollectionBook, err error) {
	item = &models.CollectionBook{}
	err = booksInCollectionQuery(db, collection).
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
	db dbx.Builder,
	collection *models.Collection,
) *dbx.SelectQuery {
	return models.CollectionMemberQuery(db).
		AndWhere(dbx.HashExp{
			"collection": collection.Id,
		})
}

func fetchMembersInCollection(
	db dbx.Builder,
	collection *models.Collection,
	page uint,
	perPage int,
) (items []*models.CollectionMember, err error) {
	items = []*models.CollectionMember{}
	err = membersInCollectionQuery(db, collection).
		Limit(int64(perPage)).
		Offset(int64(page-1) * int64(perPage)).
		All(&items)
	return
}

func countMembersInCollection(
	db dbx.Builder,
	collection *models.Collection,
) (count uint, err error) {
	type countData struct {
		Count uint `db:"count"`
	}
	result := &countData{
		Count: 0,
	}
	err = membersInCollectionQuery(db, collection).
		Select("COUNT(id) AS count").
		One(&result)
	return result.Count, err
}

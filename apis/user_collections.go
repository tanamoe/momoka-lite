package apis

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	pmodels "github.com/pocketbase/pocketbase/models"
	"tana.moe/momoka-lite/models"
)

func registerUserCollectionsRoute(
	app *pocketbase.PocketBase,
	core *core.ServeEvent,
) error {
	core.Router.GET(
		"/api/user-collections",
		listRouteHandler(app, core, onRequestUserCollections),
		apis.ActivityLogger(app),
	)
	core.Router.GET(
		"/api/user-collections/:userId",
		listRouteHandler(app, core, onRequestUserCollections),
		apis.ActivityLogger(app),
	)
	return nil
}

func onRequestUserCollections(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []*models.CollectionMember, rpage uint, rperPage int, totalItems uint, totalPages uint, err error) {
	admin, _ := c.Get(apis.ContextAdminKey).(*pmodels.Admin)
	record, _ := c.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
	userId := c.PathParam("userId")
	if userId == "" {
		if record != nil {
			userId = record.Id
		}
	}
	if userId == "" {
		return nil, 1, -1, 0, 1, notFoundError
	}
	if admin != nil {
		items, err = fetchUserJoinedCollections(app.Dao(), userId, false)
	} else {
		if record == nil {
			items, err = fetchUserJoinedCollections(app.Dao(), userId, true)
		} else if record.Id == userId {
			items, err = fetchUserJoinedCollections(app.Dao(), userId, false)
		} else {
			items, err = fetchCommonJoinedCollections(app.Dao(), userId, record.Id)
		}
	}
	if err != nil {
		return nil, 1, -1, 0, 1, err
	}
	for _, item := range items {
		if err := item.Expand(app.Dao(), expand); err != nil {
			return nil, 1, -1, 0, 1, err
		}
	}
	return items, 1, -1, uint(len(items)), 1, nil
}

func fetchUserJoinedCollections(
	dao *daos.Dao,
	userId string,
	publicOnly bool,
) (items []*models.CollectionMember, err error) {
	collectionTableName := (&models.Collection{}).TableName()
	collectionMemberTableName := (&models.CollectionMember{}).TableName()

	query := models.CollectionMemberQuery(dao).
		LeftJoin(
			collectionTableName,
			dbx.NewExp(
				fmt.Sprintf(
					"%s.collection = %s.id",
					collectionMemberTableName,
					collectionTableName,
				),
			),
		)
	if publicOnly {
		query = query.Where(
			dbx.HashExp{
				fmt.Sprintf("%s.visibility", collectionTableName): models.CollectionPublic,
			},
		)
	}
	err = query.All(&items)
	return
}

func fetchCommonJoinedCollections(
	dao *daos.Dao,
	targetId string,
	viewerId string,
) (items []*models.CollectionMember, err error) {
	return nil, unimplementedError
}

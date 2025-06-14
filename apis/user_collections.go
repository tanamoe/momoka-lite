package apis

import (
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/tools"
)

func registerUserCollectionsRoute(
	app *pocketbase.PocketBase,
	core *core.ServeEvent,
) error {
	core.Router.GET(
		"/api/user-collections",
		listRouteHandler(app, core, onRequestUserCollections),
	)
	core.Router.GET(
		"/api/user-collections/{userId}",
		listRouteHandler(app, core, onRequestUserCollections),
	)
	return nil
}

func onRequestUserCollections(
	app *pocketbase.PocketBase,
	e *core.RequestEvent,
	page uint,
	perPage int,
	expand models.ExpandMap,
) (items []*models.CollectionMember, rpage uint, rperPage int, totalItems uint, totalPages uint, err error) {
	info, err := e.RequestInfo()
	if err != nil {
		return nil, 1, -1, 0, 1, err
	}
	userId := e.Request.PathValue("userId")
	if userId == "" {
		if info.Auth != nil {
			userId = info.Auth.Id
		}
	}
	if userId == "" {
		return nil, 1, -1, 0, 1, notFoundError
	}
	if (info.Auth != nil) && (info.Auth.IsSuperuser()) {
		items, err = fetchUserJoinedCollections(app.DB(), userId, false)
	} else {
		if info.Auth == nil {
			items, err = fetchUserJoinedCollections(app.DB(), userId, true)
		} else if info.Auth.Id == userId {
			items, err = fetchUserJoinedCollections(app.DB(), userId, false)
		} else {
			items, err = fetchCommonJoinedCollections(app.DB(), userId, info.Auth.Id)
		}
	}
	if err != nil {
		return nil, 1, -1, 0, 1, err
	}
	for _, item := range items {
		if err := item.Expand(app.DB(), expand); err != nil {
			return nil, 1, -1, 0, 1, err
		}
	}
	return items, 1, -1, uint(len(items)), 1, nil
}

func fetchUserJoinedCollections(
	db dbx.Builder,
	userId string,
	publicOnly bool,
) (items []*models.CollectionMember, err error) {
	collectionTableName := (&models.Collection{}).TableName()
	collectionMembersTableName := (&models.CollectionMember{}).TableName()

	query := models.CollectionMemberQuery(db).
		LeftJoin(
			collectionTableName,
			dbx.NewExp(
				fmt.Sprintf(
					"%s.collection = %s.id",
					collectionMembersTableName,
					collectionTableName,
				),
			),
		).
		Where(
			dbx.HashExp{
				fmt.Sprintf("%s.user", collectionMembersTableName): userId,
			},
		)
	if publicOnly {
		query = query.Where(
			dbx.HashExp{
				fmt.Sprintf("%s.user", collectionMembersTableName): userId,
				fmt.Sprintf("%s.visibility", collectionTableName):  models.CollectionPublic,
			},
		)
	}
	err = query.All(&items)
	return
}

func fetchCommonJoinedCollections(
	db dbx.Builder,
	targetId string,
	viewerId string,
) (items []*models.CollectionMember, err error) {
	collectionTableName := (&models.Collection{}).TableName()
	collectionMembersTableName := (&models.CollectionMember{}).TableName()

	weightColumn := fmt.Sprintf(
		"(CASE WHEN %s.user='%s' THEN 2 "+
			"WHEN %s.visibility='%s' THEN 3 "+
			"ELSE 1 END) "+
			"AS weight",
		collectionMembersTableName, tools.EscapeSql(viewerId),
		collectionTableName, models.CollectionPublic,
	)
	scoreColumn := "SUM(weight)"
	minScore := 3

	err = models.CollectionMemberQuery(db).
		Select(
			fmt.Sprintf("%s.*", collectionMembersTableName),
			weightColumn,
		).
		LeftJoin(
			collectionTableName,
			dbx.NewExp(
				fmt.Sprintf(
					"%s.collection = %s.id",
					collectionMembersTableName,
					collectionTableName,
				),
			),
		).
		Where(
			dbx.HashExp{
				fmt.Sprintf("%s.user", collectionMembersTableName): []any{
					targetId,
					viewerId,
				},
			},
		).
		GroupBy(fmt.Sprintf("%s.id", collectionTableName)).
		Having(
			dbx.NewExp(
				fmt.Sprintf("%s >= {:minScore}", scoreColumn),
				dbx.Params{
					"minScore": minScore,
				},
			),
		).
		All(&items)
	return
}

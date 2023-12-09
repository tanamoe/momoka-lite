package hooks

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	pmodels "github.com/pocketbase/pocketbase/models"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/tools"
)

func registerAppendInCollectionsMetadataHook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	targetCollections := []string{"books", "bookDetails"}
	app.OnRecordViewRequest(targetCollections...).Add(func(e *core.RecordViewEvent) error {
		user, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
		if user == nil {
			return nil
		}
		return appendInCollectionsMetadata(
			app,
			context,
			e.HttpContext,
			user,
			[](*pmodels.Record){e.Record},
		)
	})

	app.OnRecordsListRequest(targetCollections...).Add(func(e *core.RecordsListEvent) error {
		user, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*pmodels.Record)
		if user == nil {
			return nil
		}
		return appendInCollectionsMetadata(
			app,
			context,
			e.HttpContext,
			user,
			e.Records,
		)
	})

	return nil
}

func appendInCollectionsMetadata(
	app *pocketbase.PocketBase,
	context *models.AppContext,
	c echo.Context,
	user *pmodels.Record,
	records []*pmodels.Record,
) error {
	expand, err := tools.ExtractExpandMap(c)
	if err != nil {
		return err
	}
	if _, exist := expand["metadata"]; exist {
		expand = expand["metadata"]
	} else {
		expand = models.ExpandMap{}
	}
	if _, exist := expand["inCollections"]; exist {
		expand = expand["inCollections"]
	}
	bookToCollectionsMap, err := booksWithBelongCollectionsMap(
		app.Dao(),
		user,
		records,
	)
	if err != nil {
		return err
	}
	for _, record := range records {
		collections, exist := bookToCollectionsMap[record.Id]
		if !exist {
			continue
		}
		for _, collection := range collections {
			if err := collection.Expand(app.Dao(), expand); err != nil {
				return err
			}
		}
		appendMetadata(
			record,
			map[string]interface{}{
				"inCollections": collections,
			},
		)
	}
	return nil
}

func booksWithBelongCollectionsMap(
	dao *daos.Dao,
	user *pmodels.Record,
	records []*pmodels.Record,
) (map[string][]*models.Collection, error) {
	bookCollections := []struct {
		*models.Collection

		BookId string `db:"bookId"`
	}{}

	bookIds := []any{}
	for _, record := range records {
		bookIds = append(bookIds, record.Id)
	}

	collectionTable := (&models.Collection{}).TableName()
	collectionMemberTable := (&models.CollectionMember{}).TableName()
	collectionBookTable := (&models.CollectionBook{}).TableName()

	collectionColumns := fmt.Sprintf("%s.*", collectionTable)
	bookIdColumn := fmt.Sprintf("%s.book as bookId", collectionBookTable)

	err := models.CollectionBookQuery(dao).
		Select(collectionColumns, bookIdColumn).
		RightJoin(
			collectionMemberTable,
			dbx.And(
				dbx.NewExp(
					fmt.Sprintf(
						"%s.collection = %s.collection",
						collectionBookTable,
						collectionMemberTable,
					),
				),
				dbx.HashExp{
					fmt.Sprintf("%s.user", collectionMemberTable): user.Id,
					fmt.Sprintf("%s.role", collectionMemberTable): []any{
						models.CollectionEditorRole,
					},
				},
			),
		).
		RightJoin(
			collectionTable,
			dbx.NewExp(
				fmt.Sprintf(
					"%s.collection = %s.id",
					collectionMemberTable,
					collectionTable,
				),
			),
		).
		Where(
			dbx.HashExp{
				fmt.Sprintf("%s.book", collectionBookTable): bookIds,
			},
		).
		All(&bookCollections)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	collectionMapping := map[string][]*models.Collection{}
	for _, bookCollection := range bookCollections {
		collectionMapping[bookCollection.BookId] = append(
			collectionMapping[bookCollection.BookId],
			bookCollection.Collection,
		)
	}
	return collectionMapping, nil
}

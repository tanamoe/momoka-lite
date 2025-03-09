package hooks

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/tools"
)

func registerAppendInCollectionsMetadataHook(
	app *pocketbase.PocketBase,
) error {
	targetCollections := []string{"books", "bookDetails"}
	app.OnRecordViewRequest(targetCollections...).BindFunc(func(e *core.RecordRequestEvent) error {
		info, err := e.RequestInfo()
		if err != nil {
			return err
		}
		if info.Auth == nil {
			return e.Next()
		}
		if info.Auth.IsSuperuser() {
			return e.Next()
		}
		user, err := models.FindUserById(app.DB(), info.Auth.Id)
		if err != nil {
			return err
		}
		err = appendInCollectionsMetadata(
			app,
			info,
			user,
			[](*core.Record){e.Record},
		)
		if err != nil {
			return err
		}
		return e.Next()
	})

	app.OnRecordsListRequest(targetCollections...).BindFunc(func(e *core.RecordsListRequestEvent) error {
		info, err := e.RequestInfo()
		if err != nil {
			return err
		}
		if info.Auth == nil {
			return e.Next()
		}
		if info.Auth.IsSuperuser() {
			return e.Next()
		}
		user, err := models.FindUserById(app.DB(), info.Auth.Id)
		if err != nil {
			return err
		}
		err = appendInCollectionsMetadata(
			app,
			info,
			user,
			e.Records,
		)
		if err != nil {
			return err
		}
		return e.Next()
	})

	return nil
}

func appendInCollectionsMetadata(
	app *pocketbase.PocketBase,
	c *core.RequestInfo,
	user *models.User,
	records []*core.Record,
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
		app.DB(),
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
			if err := collection.Expand(app.DB(), expand); err != nil {
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
	db dbx.Builder,
	user *models.User,
	records []*core.Record,
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

	err := models.CollectionBookQuery(db).
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

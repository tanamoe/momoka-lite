package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*CollectionBook)(nil)

type BookReadingStatus = string

const (
	BookPlanning  BookReadingStatus = "PLANNING"
	BookCompleted BookReadingStatus = "COMPLETED"
)

type CollectionBook struct {
	models.BaseModel

	CollectionId string            `db:"collection" json:"collectionId"`
	BookId       string            `db:"book" json:"bookId"`
	Quantity     int               `db:"quantity" json:"quantity"`
	Status       BookReadingStatus `db:"status" json:"status"`
}

func (m *CollectionBook) TableName() string {
	return "collectionBooks"
}

func CollectionBookQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&CollectionBook{})
}

func FindCollectionBookById(dao *daos.Dao, id string) (*CollectionBook, error) {
	collectionBook := &CollectionBook{}
	err := CollectionQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(collectionBook)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return collectionBook, nil
}

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
	Collection   *Collection       `json:"collection,omitempty"`
	BookId       string            `db:"book" json:"bookId"`
	Book         *Book             `json:"book,omitempty"`
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

func (m *CollectionBook) Expand(dao *daos.Dao, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["collection"]; exist {
		collection, err := FindCollectionById(dao, m.CollectionId)
		if err != nil {
			return err
		}
		if collection != nil {
			if err := collection.Expand(dao, e["collection"]); err != nil {
				return err
			}
		}
		m.Collection = collection
	}

	if _, exist := e["book"]; exist {
		book, err := FindBookById(dao, m.BookId)
		if err != nil {
			return err
		}
		if book != nil {
			if err := book.Expand(dao, e["book"]); err != nil {
				return err
			}
		}
		m.Book = book
	}

	return nil
}

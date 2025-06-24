package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
)

type BookReadingStatus = string

const (
	BookPlanning  BookReadingStatus = "PLANNING"
	BookCompleted BookReadingStatus = "COMPLETED"
)

type CollectionBook struct {
	Id           string            `db:"id" json:"id"`
	CollectionId string            `db:"collection" json:"collectionId"`
	Collection   *Collection       `db:"-" json:"collection,omitempty"`
	BookId       string            `db:"book" json:"bookId"`
	Book         *Book             `db:"-" json:"book,omitempty"`
	Quantity     int               `db:"quantity" json:"quantity"`
	Status       BookReadingStatus `db:"status" json:"status"`
	Notes        string            `db:"notes" json:"notes"`
}

func (m *CollectionBook) TableName() string {
	return "collectionBooks"
}

func CollectionBookQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&CollectionBook{}).TableName())
}

func FindCollectionBookById(db dbx.Builder, id string) (*CollectionBook, error) {
	collectionBook := &CollectionBook{}
	err := CollectionBookQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(collectionBook)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return collectionBook, nil
}

func (m *CollectionBook) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["collection"]; exist {
		collection, err := FindCollectionById(db, m.CollectionId)
		if err != nil {
			return err
		}
		if collection != nil {
			if err := collection.Expand(db, e["collection"]); err != nil {
				return err
			}
		}
		m.Collection = collection
	}

	if _, exist := e["book"]; exist {
		book, err := FindBookById(db, m.BookId)
		if err != nil {
			return err
		}
		if book != nil {
			if err := book.Expand(db, e["book"]); err != nil {
				return err
			}
		}
		m.Book = book
	}

	return nil
}

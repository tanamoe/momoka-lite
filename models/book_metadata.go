package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
)

type BookMetadata struct {
	Id        string  `db:"id" json:"id"`
	BookID    string  `db:"book" json:"bookId"`
	Book      *Book   `db:"-" json:"book,omitempty"`
	ISBN      string  `db:"isbn" json:"isbn"`
	FahasaSKU string  `db:"fahasaSKU" json:"fahasaSKU"`
	SizeX     float64 `db:"sizeX" json:"sizeX"`
	SizeY     float64 `db:"sizeY" json:"sizeY"`
	SizeZ     float64 `db:"sizeZ" json:"sizeZ"`
	PageCount int     `db:"pageCount" json:"pageCount"`
	Weight    float64 `db:"weight" json:"weight"`
}

func (m *BookMetadata) TableName() string {
	return "bookMetadata"
}

func BookMetadataQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&BookMetadata{}).TableName())
}

func FindBookMetadataById(db dbx.Builder, id string) (*BookMetadata, error) {
	book := &BookMetadata{}
	err := BookMetadataQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(book)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (m *BookMetadata) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["book"]; exist {
		book, err := FindBookById(db, m.BookID)
		if err != nil {
			return err
		}
		if book != nil {
			if err := book.Expand(db, e["book"]); err != nil {
				return err
			}
			m.Book = book
		}
	}

	return nil
}

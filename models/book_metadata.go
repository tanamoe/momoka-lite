package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*BookMetadata)(nil)

type BookMetadata struct {
	models.BaseModel

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

func BookMetadataQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&BookMetadata{})
}

func FindBookMetadataById(dao *daos.Dao, id string) (*BookMetadata, error) {
	book := &BookMetadata{}
	err := BookMetadataQuery(dao).
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

func (m *BookMetadata) Expand(dao *daos.Dao, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["book"]; exist {
		book, err := FindBookById(dao, m.BookID)
		if err != nil {
			return err
		}
		if book != nil {
			if err := book.Expand(dao, e["book"]); err != nil {
				return err
			}
			m.Book = book
		}
	}

	return nil
}

package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*Book)(nil)

type Book struct {
	models.BaseModel

	PublicationID    string                  `db:"publication" json:"publicationId"`
	Publication      *Publication            `db:"-" json:"publication,omitempty"`
	Edition          string                  `db:"edition" json:"edition"`
	PublishDate      types.DateTime          `db:"publishDate" json:"publishDate"`
	Covers           types.JsonArray[string] `db:"covers" json:"covers"`
	Price            int                     `db:"price" json:"price"`
	Note             string                  `db:"note" json:"note"`
	Metadata         types.JsonMap           `db:"metadata" json:"metadata"`
	ParentID         string                  `db:"parentId" json:"parentId"`
	ParentCollection string                  `db:"parentCollection" json:"parentCollection"`
}

func (m *Book) TableName() string {
	return "bookDetails"
}

func BookQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Book{})
}

func FindBookById(dao *daos.Dao, id string) (*Book, error) {
	book := &Book{}
	err := BookQuery(dao).
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

func (m *Book) Expand(dao *daos.Dao, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["publication"]; exist {
		publication, err := FindPublicationById(dao, m.PublicationID)
		if err != nil {
			return err
		}
		if publication != nil {
			if err := publication.Expand(dao, e["publication"]); err != nil {
				return err
			}
			m.Publication = publication
		}
	}

	return nil
}

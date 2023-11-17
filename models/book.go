package models

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*Book)(nil)

type Book struct {
	models.BaseModel

	PublicationID string                  `db:"publication" json:"publication"`
	Edition       string                  `db:"edition" json:"edition"`
	PublishDate   types.DateTime          `db:"publishDate" json:"publishDate"`
	Covers        types.JsonArray[string] `db:"covers" json:"covers"`
	Price         int                     `db:"price" json:"price"`
	Note          string                  `db:"note" json:"note"`
	Metadata      types.JsonMap           `db:"metadata" json:"metadata"`
}

func (m *Book) TableName() string {
	return "books"
}

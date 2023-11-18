package models

import (
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

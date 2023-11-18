package models

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*Collection)(nil)

type CollectionVisibility = string

const (
	CollectionPublic   CollectionVisibility = "PUBLIC"
	CollectionUnlisted CollectionVisibility = "UNLISTED"
	CollectionPrivate  CollectionVisibility = "PRIVATE"
)

type Collection struct {
	models.BaseModel

	OwnerId     string               `db:"publication" json:"publication"`
	Visibility  CollectionVisibility `db:"visibility" json:"visibility"`
	PublishDate types.DateTime       `db:"publishDate" json:"publishDate"`
	Name        string               `db:"name" json:"name"`
	Default     bool                 `db:"default" json:"default"`
	Description string               `db:"description" json:"description"`
	Order       string               `db:"order" json:"order"`
}

func (m *Collection) TableName() string {
	return "collections"
}

package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
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
	Owner       *User                `db:"owner,empty"`
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

func CollectionQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Collection{})
}

func FindCollectionById(dao *daos.Dao, id string) (*Collection, error) {
	collection := &Collection{}
	err := CollectionQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(collection)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (m *Collection) Expand(dao *daos.Dao, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["owner"]; exist {
		owner, err := FindUserById(dao, m.OwnerId)
		if err != nil {
			return err
		}
		m.Owner = owner
	}

	return nil
}

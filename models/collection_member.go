package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*CollectionMember)(nil)

type CollectionAccessRole = string

const (
	CollectionEditorRole CollectionAccessRole = "EDITOR"
	CollectionMemberRole CollectionAccessRole = "MEMBER"
)

type CollectionMember struct {
	models.BaseModel

	CollectionId string               `db:"collection" json:"collectionId"`
	UserId       string               `db:"user" json:"userId"`
	Role         CollectionAccessRole `db:"role" json:"role"`
}

func (m *CollectionMember) TableName() string {
	return "collectionMembers"
}

func CollectionMemberQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&CollectionMember{})
}

func FindCollectionMemberById(dao *daos.Dao, id string) (*CollectionMember, error) {
	collectionMember := &CollectionMember{}
	err := CollectionMemberQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(collectionMember)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return collectionMember, nil
}

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
	Collection   *Collection          `json:"collection,omitempty"`
	UserId       string               `db:"user" json:"userId"`
	User         *User                `json:"user,omitempty"`
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

func (m *CollectionMember) Expand(dao *daos.Dao, e ExpandMap) error {
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

	if _, exist := e["user"]; exist {
		user, err := FindUserById(dao, m.UserId)
		if err != nil {
			return err
		}
		if user != nil {
			if err := user.Expand(dao, e["user"]); err != nil {
				return err
			}
		}
		m.User = user
	}

	return nil
}

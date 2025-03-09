package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
)

type CollectionAccessRole = string

const (
	CollectionEditorRole CollectionAccessRole = "EDITOR"
	CollectionMemberRole CollectionAccessRole = "MEMBER"
)

type CollectionMember struct {
	Id           string               `db:"id" json:"id"`
	CollectionId string               `db:"collection" json:"collectionId"`
	Collection   *Collection          `db:"-" json:"collection,omitempty"`
	UserId       string               `db:"user" json:"userId"`
	User         *User                `db:"-" json:"user,omitempty"`
	Role         CollectionAccessRole `db:"role" json:"role"`
}

func (m *CollectionMember) TableName() string {
	return "collectionMembers"
}

func CollectionMemberQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&CollectionMember{}).TableName())
}

func FindCollectionMemberById(db dbx.Builder, id string) (*CollectionMember, error) {
	collectionMember := &CollectionMember{}
	err := CollectionMemberQuery(db).
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

func (m *CollectionMember) Expand(db dbx.Builder, e ExpandMap) error {
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

	if _, exist := e["user"]; exist {
		user, err := FindUserById(db, m.UserId)
		if err != nil {
			return err
		}
		if user != nil {
			if err := user.Expand(db, e["user"]); err != nil {
				return err
			}
		}
		m.User = user
	}

	return nil
}

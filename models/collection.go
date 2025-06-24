package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type CollectionVisibility = string

const (
	CollectionPublic   CollectionVisibility = "PUBLIC"
	CollectionUnlisted CollectionVisibility = "UNLISTED"
	CollectionPrivate  CollectionVisibility = "PRIVATE"
)

type Collection struct {
	Id          string               `db:"id" json:"id"`
	OwnerId     string               `db:"owner" json:"ownerId"`
	Owner       *User                `db:"-" json:"owner,omitempty"`
	Visibility  CollectionVisibility `db:"visibility" json:"visibility"`
	Name        string               `db:"name" json:"name"`
	Default     bool                 `db:"default" json:"default"`
	Description string               `db:"description" json:"description"`
	Order       int                  `db:"order" json:"order"`
}

func (m *Collection) TableName() string {
	return "collections"
}

func CollectionQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Collection{}).TableName())
}

func FindCollectionById(db dbx.Builder, id string) (*Collection, error) {
	collection := &Collection{}
	err := CollectionQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(collection)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (m *Collection) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["owner"]; exist {
		owner, err := FindUserById(db, m.OwnerId)
		if err != nil {
			return err
		}
		if owner != nil {
			if err := owner.Expand(db, e["owner"]); err != nil {
				return err
			}
			m.Owner = owner
		}
	}

	return nil
}

func FindUserDefaultCollection(db dbx.Builder, userId string) (*Collection, error) {
	collection := &Collection{}
	err := CollectionQuery(db).
		AndWhere(dbx.HashExp{
			"owner":   userId,
			"default": true,
		}).
		Limit(1).
		One(collection)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (m *Collection) userHadRole(db dbx.Builder, userId string, roles ...CollectionAccessRole) (bool, error) {
	type countData struct {
		Count uint `db:"count"`
	}
	count := &countData{
		Count: 0,
	}
	var rolesAsAny []any
	for _, role := range roles {
		rolesAsAny = append(rolesAsAny, role)
	}
	err := db.Select("*").
		Select("COUNT(id) AS count").
		AndWhere(dbx.HashExp{
			"collection": m.Id,
			"user":       userId,
			"role":       rolesAsAny,
		}).
		One(&count)
	return count.Count > 0, err
}

func (m *Collection) CanBeAccessedBy(db dbx.Builder, userId string) (bool, error) {
	if m.OwnerId == userId {
		return true, nil
	}
	if m.Visibility != CollectionPrivate {
		return true, nil
	}
	return m.userHadRole(db, userId, CollectionMemberRole, CollectionEditorRole)
}

func (m *Collection) CanBeEditedBy(db dbx.Builder, userId string) (bool, error) {
	if m.OwnerId == userId {
		return true, nil
	}
	return m.userHadRole(db, userId, CollectionEditorRole)
}

func (m *Collection) HadMember(db dbx.Builder, userId string) (bool, error) {
	if m.OwnerId == userId {
		return true, nil
	}
	return m.userHadRole(db, userId, CollectionMemberRole, CollectionEditorRole)
}

func (m *Collection) AddMember(db dbx.Builder, userId string, role CollectionAccessRole) error {
	member := &CollectionMember{
		Id:           core.GenerateDefaultRandomId(),
		CollectionId: m.Id,
		UserId:       userId,
		Role:         role,
	}
	existMember := &CollectionMember{}
	err := CollectionMemberQuery(db).
		AndWhere(dbx.HashExp{
			"collection": m.Id,
			"user":       userId,
		}).
		One(&existMember)
	if (err != nil) && (!errors.Is(err, sql.ErrNoRows)) {
		return err
	}
	if err == nil {
		member = existMember
		member.Role = role
	}
	if err := db.Model(member).Insert(); err != nil {
		return err
	}
	return nil
}

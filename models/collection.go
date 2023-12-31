package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
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
		if owner != nil {
			if err := owner.Expand(dao, e["owner"]); err != nil {
				return err
			}
			m.Owner = owner
		}
	}

	return nil
}

func FindUserDefaultCollection(dao *daos.Dao, userId string) (*Collection, error) {
	collection := &Collection{}
	err := CollectionQuery(dao).
		AndWhere(dbx.HashExp{
			"owner":   userId,
			"default": true,
		}).
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

func (m *Collection) userHadRole(dao *daos.Dao, userId string, roles ...CollectionAccessRole) (bool, error) {
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
	err := CollectionMemberQuery(dao).
		Select("COUNT(id) AS count").
		AndWhere(dbx.HashExp{
			"collection": m.Id,
			"user":       userId,
			"role":       rolesAsAny,
		}).
		One(&count)
	return count.Count > 0, err
}

func (m *Collection) CanBeAccessedBy(dao *daos.Dao, userId string) (bool, error) {
	if m.OwnerId == userId {
		return true, nil
	}
	if m.Visibility != CollectionPrivate {
		return true, nil
	}
	return m.userHadRole(dao, userId, CollectionMemberRole, CollectionEditorRole)
}

func (m *Collection) CanBeEditedBy(dao *daos.Dao, userId string) (bool, error) {
	if m.OwnerId == userId {
		return true, nil
	}
	return m.userHadRole(dao, userId, CollectionEditorRole)
}

func (m *Collection) HadMember(dao *daos.Dao, userId string) (bool, error) {
	if m.OwnerId == userId {
		return true, nil
	}
	return m.userHadRole(dao, userId, CollectionMemberRole, CollectionEditorRole)
}

func (m *Collection) AddMember(dao *daos.Dao, userId string, role CollectionAccessRole) error {
	member := &CollectionMember{
		CollectionId: m.Id,
		UserId:       userId,
		Role:         role,
	}
	existMember := &CollectionMember{}
	err := CollectionMemberQuery(dao).
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
	if err := dao.Save(member); err != nil {
		return err
	}
	return nil
}

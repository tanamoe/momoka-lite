package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
)

type User struct {
	Id          string `db:"id" json:"id"`
	Username    string `db:"username" json:"username"`
	Email       string `db:"email" json:"-"`
	DisplayName string `db:"displayName" json:"displayName"`
	Bio         string `db:"bio" json:"bio"`
	Avatar      string `db:"avatar" json:"avatar"`
	Banner      string `db:"banner" json:"banner"`
}

func (m *User) TableName() string {
	return "users"
}

func UserQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&User{}).TableName())
}

func FindUserById(db dbx.Builder, id string) (*User, error) {
	user := &User{}
	err := UserQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(user)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *User) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

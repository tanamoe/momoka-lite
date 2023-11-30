package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*User)(nil)

type User struct {
	models.BaseModel

	Username    string `db:"username" json:"username"`
	Email       string `db:"email" json:"email"`
	DisplayName string `db:"displayName" json:"displayName"`
	Bio         string `db:"bio" json:"bio"`
	Avatar      string `db:"avatar" json:"avatar"`
	Banner      string `db:"banner" json:"banner"`
}

func (m *User) TableName() string {
	return "users"
}

func UserQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&User{})
}

func FindUserById(dao *daos.Dao, id string) (*User, error) {
	user := &User{}
	err := UserQuery(dao).
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

func (m *User) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

package models

import (
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

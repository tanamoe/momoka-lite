package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Staff)(nil)

type Staff struct {
	models.BaseModel

	Name string `db:"name" json:"name"`
}

func (m *Staff) TableName() string {
	return "staffs"
}

func StaffQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Staff{})
}

func FindStaffById(dao *daos.Dao, id string) (*Staff, error) {
	staff := &Staff{}
	err := GenreQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(staff)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return staff, nil
}

func (m *Staff) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

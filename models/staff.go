package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
)

type Staff struct {
	Id   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (m *Staff) TableName() string {
	return "staffs"
}

func StaffQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Staff{}).TableName())
}

func FindStaffById(db dbx.Builder, id string) (*Staff, error) {
	staff := &Staff{}
	err := GenreQuery(db).
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

func (m *Staff) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

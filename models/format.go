package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Format)(nil)

type Format struct {
	models.BaseModel

	Name        string `db:"name" json:"name"`
	Slug        string `db:"slug" json:"slug"`
	Color       string `db:"color" json:"color"`
	Description string `db:"description" json:"description"`
	Thumbnail   string `db:"thumbnail" json:"thumbnail"`
}

func (m *Format) TableName() string {
	return "formats"
}

func FormatQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Format{})
}

func FindFormatById(dao *daos.Dao, id string) (*Format, error) {
	format := &Format{}
	err := FormatQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(format)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return format, nil
}

func (m *Format) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

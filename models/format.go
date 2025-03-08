package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
)

type Format struct {
	Id          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Slug        string `db:"slug" json:"slug"`
	Color       string `db:"color" json:"color"`
	Description string `db:"description" json:"description"`
	Thumbnail   string `db:"thumbnail" json:"thumbnail"`
}

func (m *Format) TableName() string {
	return "formats"
}

func FormatQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Format{}).TableName())
}

func FindFormatById(db dbx.Builder, id string) (*Format, error) {
	format := &Format{}
	err := FormatQuery(db).
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

func (m *Format) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

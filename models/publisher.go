package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
)

type Publisher struct {
	Id    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Logo  string `db:"logo" json:"logo"`
	Slug  string `db:"slug" json:"slug"`
	Color string `db:"color" json:"color"`
}

func (m *Publisher) TableName() string {
	return "publishers"
}

func PublisherQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Publisher{}).TableName())
}

func FindPublisherById(db dbx.Builder, id string) (*Publisher, error) {
	publisher := &Publisher{}
	err := PublisherQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(publisher)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

func (m *Publisher) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

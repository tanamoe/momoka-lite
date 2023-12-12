package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Publisher)(nil)

type Publisher struct {
	models.BaseModel

	Name  string `db:"name" json:"name"`
	Logo  string `db:"logo" json:"logo"`
	Slug  string `db:"slug" json:"slug"`
	Color string `db:"color" json:"color"`
}

func (m *Publisher) TableName() string {
	return "publishers"
}

func PublisherQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Publisher{})
}

func FindPublisherById(dao *daos.Dao, id string) (*Publisher, error) {
	publisher := &Publisher{}
	err := PublisherQuery(dao).
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

func (m *Publisher) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
)

type Demographic struct {
	Id   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`
}

func (m *Demographic) TableName() string {
	return "demographics"
}

func DemographicQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Demographic{}).TableName())
}

func FindDemographicById(db dbx.Builder, id string) (*Demographic, error) {
	demographic := &Demographic{}
	err := DemographicQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(demographic)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return demographic, nil
}

func (m *Demographic) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

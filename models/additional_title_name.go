package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
)

type AdditionalTitleName struct {
	Id       string `db:"id" json:"id"`
	TitleId  string `db:"title" json:"titleId"`
	Title    *Title `db:"-" json:"title,omitempty"`
	Language string `db:"language" json:"language"`
	Name     string `db:"name" json:"name"`
}

func (m *AdditionalTitleName) TableName() string {
	return "additionalTitleNames"
}

func AdditionalTitleNameQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&AdditionalTitleName{}).TableName())
}

func FindAdditionalTitleNameById(db dbx.Builder, id string) (*AdditionalTitleName, error) {
	name := &AdditionalTitleName{}
	err := AdditionalTitleNameQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return name, nil
}

func (m *AdditionalTitleName) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["title"]; exist {
		title, err := FindTitleById(db, m.TitleId)
		if err != nil {
			return err
		}
		if title != nil {
			if err := title.Expand(db, e["title"]); err != nil {
				return err
			}
			m.Title = title
		}
	}

	return nil
}

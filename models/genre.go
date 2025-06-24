package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
)

type Genre struct {
	Id   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`
}

func (m *Genre) TableName() string {
	return "genres"
}

func GenreQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Genre{}).TableName())
}

func FindGenreById(db dbx.Builder, id string) (*Genre, error) {
	genre := &Genre{}
	err := GenreQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(genre)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return genre, nil
}

func (m *Genre) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

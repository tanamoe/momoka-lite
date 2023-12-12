package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Genre)(nil)

type Genre struct {
	models.BaseModel

	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`
}

func (m *Genre) TableName() string {
	return "genres"
}

func GenreQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Genre{})
}

func FindGenreById(dao *daos.Dao, id string) (*Genre, error) {
	genre := &Genre{}
	err := GenreQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(genre)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return genre, nil
}

func (m *Genre) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

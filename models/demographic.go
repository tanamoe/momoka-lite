package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Demographic)(nil)

type Demographic struct {
	models.BaseModel

	Name string `db:"name" json:"name"`
	Slug string `db:"slug" json:"slug"`
}

func (m *Demographic) TableName() string {
	return "demographics"
}

func DemographicQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Demographic{})
}

func FindDemographicById(dao *daos.Dao, id string) (*Demographic, error) {
	demographic := &Demographic{}
	err := DemographicQuery(dao).
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

func (m *Demographic) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

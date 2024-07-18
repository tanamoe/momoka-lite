package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*AdditionalTitleName)(nil)

type AdditionalTitleName struct {
	models.BaseModel

	TitleId  string `db:"title" json:"titleId"`
	Title    *Title `db:"-" json:"title,omitempty"`
	Language string `db:"language" json:"language"`
	Name     string `db:"name" json:"name"`
}

func (m *AdditionalTitleName) TableName() string {
	return "additionalTitleNames"
}

func AdditionalTitleNameQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&AdditionalTitleName{})
}

func FindAdditionalTitleNameById(dao *daos.Dao, id string) (*AdditionalTitleName, error) {
	name := &AdditionalTitleName{}
	err := AdditionalTitleNameQuery(dao).
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

func (m *AdditionalTitleName) Expand(dao *daos.Dao, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["title"]; exist {
		title, err := FindTitleById(dao, m.TitleId)
		if err != nil {
			return err
		}
		if title != nil {
			if err := title.Expand(dao, e["title"]); err != nil {
				return err
			}
			m.Title = title
		}
	}

	return nil
}

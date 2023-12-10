package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*Title)(nil)

type Title struct {
	models.BaseModel

	Slug        string                  `db:"slug" json:"slug"`
	Name        string                  `db:"name" json:"name"`
	Description string                  `db:"description" json:"description"`
	Format      string                  `db:"format" json:"format"`
	Cover       string                  `db:"cover" json:"cover"`
	Demographic string                  `db:"demographic" json:"demographic"`
	Genres      types.JsonArray[string] `db:"genres" json:"genres"`
	Metadata    types.JsonMap           `db:"metadata" json:"metadata"`
}

func (m *Title) TableName() string {
	return "titles"
}

func TitleQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Title{})
}

func FindTitleById(dao *daos.Dao, id string) (*Title, error) {
	title := &Title{}
	err := TitleQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(title)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return title, nil
}

func (m *Title) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

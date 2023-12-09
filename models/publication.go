package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*Publication)(nil)

type Publication struct {
	models.BaseModel

	ReleaseID string                  `db:"release" json:"release"`
	Name      string                  `db:"name" json:"name"`
	Volume    int                     `db:"volume" json:"volume"`
	Covers    types.JsonArray[string] `db:"covers" json:"covers"`
	Digital   bool                    `db:"digital" json:"digital"`
	Metadata  types.JsonMap           `db:"metadata" json:"metadata"`
}

func (m *Publication) TableName() string {
	return "publications"
}

func PublicationQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Publication{})
}

func FindPublicationById(dao *daos.Dao, id string) (*Publication, error) {
	publication := &Publication{}
	err := PublicationQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(publication)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return publication, nil
}

func (m *Publication) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

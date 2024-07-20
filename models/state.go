package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*State)(nil)

const (
	ImagorSecretStateId      = "000ImagorSecret"
	AssetImageResizedStateId = "000AssetResized"
)

type State struct {
	models.BaseModel

	Value string `db:"value" json:"value"`
}

func (m *State) TableName() string {
	return "states"
}

func StateQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&State{})
}

func FindStateById(dao *daos.Dao, id string) (*State, error) {
	assetType := &State{}
	err := StateQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(assetType)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return assetType, nil
}

func (m *State) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

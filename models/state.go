package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
)

const (
	ImagorSecretStateId      = "000ImagorSecret"
	AssetImageResizedStateId = "000AssetResized"
)

type State struct {
	Id    string `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}

func (m *State) TableName() string {
	return "states"
}

func StateQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&State{}).TableName())
}

func FindStateById(db dbx.Builder, id string) (*State, error) {
	assetType := &State{}
	err := StateQuery(db).
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

func (m *State) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

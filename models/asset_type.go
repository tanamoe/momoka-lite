package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
)

const AssetTypeCoverID = "0000000000cover"

type AssetType struct {
	Id   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (m *AssetType) TableName() string {
	return "assetTypes"
}

func AssetTypeQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&AssetType{}).TableName())
}

func FindAssetTypeById(db dbx.Builder, id string) (*AssetType, error) {
	assetType := &AssetType{}
	err := AssetTypeQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(assetType)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return assetType, nil
}

func (m *AssetType) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

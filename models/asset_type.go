package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*AssetType)(nil)

const AssetTypeCoverID = "0000000000cover"

type AssetType struct {
	models.BaseModel

	Name string `db:"name" json:"name"`
}

func (m *AssetType) TableName() string {
	return "assetTypes"
}

func AssetTypeQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&AssetType{})
}

func FindAssetTypeById(dao *daos.Dao, id string) (*AssetType, error) {
	assetType := &AssetType{}
	err := AssetTypeQuery(dao).
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

func (m *AssetType) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

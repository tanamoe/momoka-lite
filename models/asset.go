package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*Asset)(nil)

type Asset struct {
	models.BaseModel

	BookId       string        `db:"book" json:"bookId"`
	Book         *Book         `db:"-" json:"book,omitempty"`
	TypeId       string        `db:"type" json:"typeId"`
	Type         *AssetType    `db:"-" json:"type,omitempty"`
	Image        string        `db:"image" json:"image"`
	ResizedImage types.JsonMap `db:"resizedImage" json:"resizedImage"`
	Priority     int           `db:"priority" json:"priority"`
}

func (m *Asset) TableName() string {
	return "assets"
}

func AssetQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Asset{})
}

func FindAssetById(dao *daos.Dao, id string) (*Asset, error) {
	asset := &Asset{}
	err := AssetQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(asset)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (m *Asset) Expand(dao *daos.Dao, e ExpandMap) error {
	return nil
}

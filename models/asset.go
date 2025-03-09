package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

type Asset struct {
	Id           string                `db:"id" json:"id"`
	BookId       string                `db:"book" json:"bookId"`
	Book         *Book                 `db:"-" json:"book,omitempty"`
	Description  string                `db:"description" json:"description"`
	TypeId       string                `db:"type" json:"typeId"`
	Type         *AssetType            `db:"-" json:"type,omitempty"`
	Image        string                `db:"image" json:"image"`
	ResizedImage types.JSONMap[string] `db:"resizedImage" json:"resizedImage"`
	Priority     int                   `db:"priority" json:"priority"`
}

func (m *Asset) TableName() string {
	return "assets"
}

func AssetQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Asset{}).TableName())
}

func FindAssetById(db dbx.Builder, id string) (*Asset, error) {
	asset := &Asset{}
	err := AssetQuery(db).
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

func (m *Asset) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

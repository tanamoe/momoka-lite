package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

type Book struct {
	Id            string                  `db:"id" json:"id"`
	PublicationID string                  `db:"publication" json:"publicationId"`
	Publication   *Publication            `db:"-" json:"publication,omitempty"`
	Edition       string                  `db:"edition" json:"edition"`
	PublishDate   types.DateTime          `db:"publishDate" json:"publishDate"`
	Covers        types.JSONArray[string] `db:"covers" json:"covers"`
	Price         int                     `db:"price" json:"price"`
	Note          string                  `db:"note" json:"note"`
	Metadata      types.JSONMap[any]      `db:"metadata" json:"metadata"`
	DefaultAsset  *Asset                  `db:"-" json:"defaultAsset,omitempty"`
}

func (m *Book) TableName() string {
	return "books"
}

func BookQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Book{}).TableName())
}

func FindBookById(db dbx.Builder, id string) (*Book, error) {
	book := &Book{}
	err := BookQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(book)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (m *Book) GetRoot(db dbx.Builder) (*Book, error) {
	publication, err := FindPublicationById(db, m.PublicationID)
	if err != nil {
		return nil, err
	}
	if m.Id == publication.DefaultBookId {
		return m, nil
	}
	book, err := FindBookById(db, publication.DefaultBookId)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (m *Book) GetDefaultAsset(db dbx.Builder) (*Asset, error) {
	asset := &Asset{}
	err := AssetQuery(db).
		AndWhere(dbx.HashExp{
			"book": m.Id,
			"type": AssetTypeCoverID,
		}).
		OrderBy("priority ASC").
		Limit(1).
		One(asset)
	if errors.Is(err, sql.ErrNoRows) {
		root, err := m.GetRoot(db)
		if root.Id == m.Id {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		return root.GetDefaultAsset(db)
	}
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (m *Book) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["publication"]; exist {
		publication, err := FindPublicationById(db, m.PublicationID)
		if err != nil {
			return err
		}
		if publication != nil {
			if err := publication.Expand(db, e["publication"]); err != nil {
				return err
			}
			m.Publication = publication
		}
	}

	if _, exist := e["defaultAsset"]; exist {
		asset, err := m.GetDefaultAsset(db)
		if err != nil {
			return err
		}
		if asset != nil {
			if err := asset.Expand(db, e["defaultAsset"]); err != nil {
				return err
			}
			m.DefaultAsset = asset
		}
	}

	return nil
}

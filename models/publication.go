package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

type Publication struct {
	Id            string                  `db:"id" json:"id"`
	ReleaseId     string                  `db:"release" json:"releaseId"`
	Release       *Release                `db:"-" json:"release,omitempty"`
	Name          string                  `db:"name" json:"name"`
	Subtitle      string                  `db:"subtitle" json:"subtitle"`
	Volume        int                     `db:"volume" json:"volume"`
	DefaultBookId string                  `db:"defaultBook" json:"defaultBookId"`
	DefaultBook   *Book                   `db:"-" json:"defaultBook,omitempty"`
	Covers        types.JSONArray[string] `db:"covers" json:"covers"`
	Metadata      types.JSONMap[any]      `db:"-" json:"-"`
}

func (m *Publication) TableName() string {
	return "publications"
}

func PublicationQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Publication{}).TableName())
}

func FindPublicationById(db dbx.Builder, id string) (*Publication, error) {
	publication := &Publication{}
	err := PublicationQuery(db).
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

func (m *Publication) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["release"]; exist {
		release, err := FindReleaseById(db, m.ReleaseId)
		if err != nil {
			return err
		}
		if release != nil {
			if err := release.Expand(db, e["release"]); err != nil {
				return err
			}
			m.Release = release
		}
	}

	if _, exist := e["defaultBook"]; exist {
		book, err := FindBookById(db, m.DefaultBookId)
		if err != nil {
			return err
		}
		if book != nil {
			if err := book.Expand(db, e["defaultBook"]); err != nil {
				return err
			}
			m.DefaultBook = book
		}
	}

	return nil
}

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

	ReleaseId     string                  `db:"release" json:"releaseId"`
	Release       *Release                `db:"-" json:"release,omitempty"`
	Name          string                  `db:"name" json:"name"`
	Volume        int                     `db:"volume" json:"volume"`
	DefaultBookId string                  `db:"defaultBook" json:"defaultBookId"`
	DefaultBook   *Book                   `db:"-" json:"defaultBook,omitempty"`
	Covers        types.JsonArray[string] `db:"covers" json:"covers"`
	Metadata      types.JsonMap           `db:"-" json:"-"`
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
	if e == nil {
		return nil
	}

	if _, exist := e["release"]; exist {
		release, err := FindReleaseById(dao, m.ReleaseId)
		if err != nil {
			return err
		}
		if release != nil {
			if err := release.Expand(dao, e["release"]); err != nil {
				return err
			}
			m.Release = release
		}
	}

	if _, exist := e["defaultBook"]; exist {
		book, err := FindBookById(dao, m.DefaultBookId)
		if err != nil {
			return err
		}
		if book != nil {
			if err := book.Expand(dao, e["defaultBook"]); err != nil {
				return err
			}
			m.DefaultBook = book
		}
	}

	return nil
}

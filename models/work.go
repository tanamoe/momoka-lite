package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
)

type Work struct {
	Id      string `db:"id" json:"id"`
	TitleId string `db:"title" json:"titleId"`
	Title   *Title `db:"-" json:"title,omitempty"`
	StaffId string `db:"staff" json:"staffId"`
	Staff   *Staff `db:"-" json:"staff,omitempty"`
}

func (m *Work) TableName() string {
	return "works"
}

func WorkQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Work{}).TableName())
}

func FindWorkById(db dbx.Builder, id string) (*Work, error) {
	work := &Work{}
	err := GenreQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(work)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return work, nil
}

func (m *Work) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["title"]; exist {
		title, err := FindTitleById(db, m.TitleId)
		if err != nil {
			return err
		}
		if title != nil {
			if err := title.Expand(db, e["title"]); err != nil {
				return err
			}
			m.Title = title
		}
	}

	if _, exist := e["staff"]; exist {
		staff, err := FindStaffById(db, m.StaffId)
		if err != nil {
			return err
		}
		if staff != nil {
			if err := staff.Expand(db, e["staff"]); err != nil {
				return err
			}
			m.Staff = staff
		}
	}
	return nil
}

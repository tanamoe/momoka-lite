package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Work)(nil)

type Work struct {
	models.BaseModel

	TitleId string `db:"title" json:"titleId"`
	Title   *Title `db:"-" json:"title,omitempty"`
	StaffId string `db:"staff" json:"staffId"`
	Staff   *Staff `db:"-" json:"staff,omitempty"`
}

func (m *Work) TableName() string {
	return "works"
}

func WorkQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Work{})
}

func FindWorkById(dao *daos.Dao, id string) (*Work, error) {
	work := &Work{}
	err := GenreQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(work)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return work, nil
}

func (m *Work) Expand(dao *daos.Dao, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["title"]; exist {
		title, err := FindTitleById(dao, m.TitleId)
		if err != nil {
			return err
		}
		if title != nil {
			if err := title.Expand(dao, e["title"]); err != nil {
				return err
			}
			m.Title = title
		}
	}

	if _, exist := e["staff"]; exist {
		staff, err := FindStaffById(dao, m.StaffId)
		if err != nil {
			return err
		}
		if staff != nil {
			if err := staff.Expand(dao, e["staff"]); err != nil {
				return err
			}
			m.Staff = staff
		}
	}
	return nil
}

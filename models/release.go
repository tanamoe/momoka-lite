package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type ReleaseStatus string

const (
	ReleaseStatusWaitingForApproval ReleaseStatus = "WAITING_FOR_APPROVAL"
	ReleaseStatusRegistered         ReleaseStatus = "REGISTERED"
	ReleaseStatusLicensed           ReleaseStatus = "LICENSED"
	ReleaseStatusOnGoing            ReleaseStatus = "ON_GOING"
	ReleaseStatusCompleted          ReleaseStatus = "COMPLETED"
	ReleaseStatusHiatus             ReleaseStatus = "HIATUS"
	ReleaseStatusCancelled          ReleaseStatus = "CANCELLED"
)

var _ models.Model = (*Release)(nil)

type Release struct {
	models.BaseModel

	TitleId     string        `db:"title" json:"titleId"`
	Title       *Title        `db:"-" json:"title,omitempty"`
	Name        string        `db:"name" json:"name"`
	PublisherId string        `db:"publisher" json:"publisherId"`
	Publisher   *Publisher    `db:"-" json:"publisher,omitempty"`
	Status      ReleaseStatus `db:"status" json:"status"`
}

func (m *Release) TableName() string {
	return "releases"
}

func ReleaseQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Release{})
}

func FindReleaseById(dao *daos.Dao, id string) (*Release, error) {
	release := &Release{}
	err := ReleaseQuery(dao).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(release)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return release, nil
}

func (m *Release) Expand(dao *daos.Dao, e ExpandMap) error {
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

	if _, exist := e["publisher"]; exist {
		publisher, err := FindPublisherById(dao, m.PublisherId)
		if err != nil {
			return err
		}
		if publisher != nil {
			if err := publisher.Expand(dao, e["publisher"]); err != nil {
				return err
			}
			m.Publisher = publisher
		}
	}

	return nil
}

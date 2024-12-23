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

	TitleId        string        `db:"title" json:"titleId"`
	Title          *Title        `db:"-" json:"title,omitempty"`
	Name           string        `db:"name" json:"name"`
	Type           string        `db:"type" json:"type"`
	Digital        bool          `db:"digital" json:"digital"`
	Disambiguation string        `db:"disambiguation" json:"disambiguation"`
	PublisherId    string        `db:"publisher" json:"publisherId"`
	Publisher      *Publisher    `db:"-" json:"publisher,omitempty"`
	PartnerId      string        `db:"partner" json:"partnerId"`
	Partner        *Publisher    `db:"-" json:"partner,omitempty"`
	Status         ReleaseStatus `db:"status" json:"status"`
	FrontId        string        `db:"front" json:"frontId"`
	Front          *Asset        `db:"-" json:"front,omitempty"`
	BannerId       string        `db:"banner" json:"bannerId"`
	Banner         *Asset        `db:"-" json:"banner,omitempty"`
	LogoId         string        `db:"logo" json:"logoId"`
	Logo           *Asset        `db:"-" json:"logo,omitempty"`
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

	if _, exist := e["partner"]; exist {
		partner, err := FindPublisherById(dao, m.PartnerId)
		if err != nil {
			return err
		}
		if partner != nil {
			if err := partner.Expand(dao, e["partner"]); err != nil {
				return err
			}
			m.Partner = partner
		}
	}

	if _, exist := e["front"]; exist {
		front, err := FindAssetById(dao, m.FrontId)
		if err != nil {
			return err
		}
		if front != nil {
			if err := front.Expand(dao, e["front"]); err != nil {
				return err
			}
			m.Front = front
		}
	}

	if _, exist := e["banner"]; exist {
		banner, err := FindAssetById(dao, m.BannerId)
		if err != nil {
			return err
		}
		if banner != nil {
			if err := banner.Expand(dao, e["banner"]); err != nil {
				return err
			}
			m.Banner = banner
		}
	}

	if _, exist := e["logo"]; exist {
		logo, err := FindAssetById(dao, m.LogoId)
		if err != nil {
			return err
		}
		if logo != nil {
			if err := logo.Expand(dao, e["logo"]); err != nil {
				return err
			}
			m.Logo = logo
		}
	}

	return nil
}

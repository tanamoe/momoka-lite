package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
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

type Release struct {
	Id             string        `db:"id" json:"id"`
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

func ReleaseQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Release{}).TableName())
}

func FindReleaseById(db dbx.Builder, id string) (*Release, error) {
	release := &Release{}
	err := ReleaseQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(release)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return release, nil
}

func (m *Release) Expand(db dbx.Builder, e ExpandMap) error {
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

	if _, exist := e["publisher"]; exist {
		publisher, err := FindPublisherById(db, m.PublisherId)
		if err != nil {
			return err
		}
		if publisher != nil {
			if err := publisher.Expand(db, e["publisher"]); err != nil {
				return err
			}
			m.Publisher = publisher
		}
	}

	if _, exist := e["partner"]; exist {
		partner, err := FindPublisherById(db, m.PartnerId)
		if err != nil {
			return err
		}
		if partner != nil {
			if err := partner.Expand(db, e["partner"]); err != nil {
				return err
			}
			m.Partner = partner
		}
	}

	if _, exist := e["front"]; exist {
		front, err := FindAssetById(db, m.FrontId)
		if err != nil {
			return err
		}
		if front != nil {
			if err := front.Expand(db, e["front"]); err != nil {
				return err
			}
			m.Front = front
		}
	}

	if _, exist := e["banner"]; exist {
		banner, err := FindAssetById(db, m.BannerId)
		if err != nil {
			return err
		}
		if banner != nil {
			if err := banner.Expand(db, e["banner"]); err != nil {
				return err
			}
			m.Banner = banner
		}
	}

	if _, exist := e["logo"]; exist {
		logo, err := FindAssetById(db, m.LogoId)
		if err != nil {
			return err
		}
		if logo != nil {
			if err := logo.Expand(db, e["logo"]); err != nil {
				return err
			}
			m.Logo = logo
		}
	}

	return nil
}

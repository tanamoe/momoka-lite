package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

type Title struct {
	Id               string                  `db:"id" json:"id"`
	SlugGroup        string                  `db:"slugGroup" json:"-"`
	Slug             string                  `db:"slug" json:"slug"`
	Name             string                  `db:"name" json:"name"`
	Description      string                  `db:"description" json:"description"`
	FormatId         string                  `db:"format" json:"formatId"`
	Format           *Format                 `db:"-" json:"format,omitempty"`
	Cover            string                  `db:"cover" json:"cover"`
	DemographicId    string                  `db:"demographic" json:"demographicId"`
	Demographic      *Demographic            `db:"-" json:"demographic,omitempty"`
	GenreIds         types.JSONArray[string] `db:"genres" json:"genreIds"`
	Genres           []*Genre                `db:"-" json:"genres,omitempty"`
	DefaultReleaseId string                  `db:"defaultRelease" json:"defaultReleaseId"`
	DefaultRelease   *Release                `db:"-" json:"defaultRelease,omitempty"`
	Metadata         types.JSONMap[any]      `db:"metadata" json:"metadata"`
}

func (m *Title) TableName() string {
	return "titles"
}

func TitleQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Title{}).TableName())
}

func FindTitleById(db dbx.Builder, id string) (*Title, error) {
	title := &Title{}
	err := TitleQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(title)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return title, nil
}

func (m *Title) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["format"]; exist {
		format, err := FindFormatById(db, m.FormatId)
		if err != nil {
			return err
		}
		if format != nil {
			if err := format.Expand(db, e["format"]); err != nil {
				return err
			}
			m.Format = format
		}
	}

	if _, exist := e["demographic"]; exist {
		demographic, err := FindDemographicById(db, m.DemographicId)
		if err != nil {
			return err
		}
		if demographic != nil {
			if err := demographic.Expand(db, e["demographic"]); err != nil {
				return err
			}
			m.Demographic = demographic
		}
	}

	if _, exist := e["genres"]; exist {
		for _, genreId := range m.GenreIds {
			genre, err := FindGenreById(db, genreId)
			if err != nil {
				return err
			}
			if genre == nil {
				continue
			}
			if err := genre.Expand(db, e["genres"]); err != nil {
				return err
			}
			m.Genres = append(m.Genres, genre)
		}
	}

	if _, exist := e["defaultRelease"]; exist {
		release, err := FindReleaseById(db, m.DefaultReleaseId)
		if err != nil {
			return err
		}
		if release != nil {
			if err := release.Expand(db, e["defaultRelease"]); err != nil {
				return err
			}
			m.DefaultRelease = release
		}
	}

	return nil
}

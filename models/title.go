package models

import (
	"database/sql"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*Title)(nil)

type Title struct {
	models.BaseModel

	SlugGroup     string                  `db:"slugGroup" json:"-"`
	Slug          string                  `db:"slug" json:"slug"`
	Name          string                  `db:"name" json:"name"`
	Description   string                  `db:"description" json:"description"`
	FormatId      string                  `db:"format" json:"formatId"`
	Format        *Format                 `db:"-" json:"format,omitempty"`
	Cover         string                  `db:"cover" json:"cover"`
	DemographicId string                  `db:"demographic" json:"demographicId"`
	Demographic   *Demographic            `db:"-" json:"demographic,omitempty"`
	GenreIds      types.JsonArray[string] `db:"genres" json:"genreIds"`
	Genres        []*Genre                `db:"-" json:"genres,omitempty"`
	Metadata      types.JsonMap           `db:"metadata" json:"metadata"`
}

func (m *Title) TableName() string {
	return "titles"
}

func TitleQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Title{})
}

func FindTitleById(dao *daos.Dao, id string) (*Title, error) {
	title := &Title{}
	err := TitleQuery(dao).
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

func (m *Title) Expand(dao *daos.Dao, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["format"]; exist {
		format, err := FindFormatById(dao, m.FormatId)
		if err != nil {
			return err
		}
		if format != nil {
			if err := format.Expand(dao, e["format"]); err != nil {
				return err
			}
			m.Format = format
		}
	}

	if _, exist := e["demographic"]; exist {
		demographic, err := FindDemographicById(dao, m.DemographicId)
		if err != nil {
			return err
		}
		if demographic != nil {
			if err := demographic.Expand(dao, e["demographic"]); err != nil {
				return err
			}
			m.Demographic = demographic
		}
	}

	if _, exist := e["genres"]; exist {
		for _, genreId := range m.GenreIds {
			genre, err := FindGenreById(dao, genreId)
			if err != nil {
				return err
			}
			if genre == nil {
				continue
			}
			if err := genre.Expand(dao, e["genres"]); err != nil {
				return err
			}
			m.Genres = append(m.Genres, genre)
		}
	}

	return nil
}

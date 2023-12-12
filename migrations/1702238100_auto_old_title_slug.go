package migrations

import (
	"database/sql"
	"errors"

	"github.com/gosimple/slug"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func titlesWithoutSlugQuery(dao *daos.Dao) *dbx.SelectQuery {
	return models.
		TitleQuery(dao).
		AndWhere(
			dbx.HashExp{
				"slug": "",
			},
		)
}

func init() {
	m.Register(func(db dbx.Builder) error {
		type countData struct {
			Count int64 `db:"count"`
		}
		dao := daos.New(db)
		count := &countData{}
		err := titlesWithoutSlugQuery(dao).
			Select("COUNT(id) AS count").
			One(&count)
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		if err != nil {
			return err
		}
		var errChans [](chan error)
		pageSize := int64(25)
		for offset := int64(0); offset < count.Count; offset += pageSize {
			currentOffset := offset
			i := len(errChans)
			errChans = append(errChans, make(chan error))
			go (func() {
				titles := []*models.Title{}
				err := titlesWithoutSlugQuery(dao).
					Offset(currentOffset).
					Limit(pageSize).
					All(&titles)
				if errors.Is(err, sql.ErrNoRows) {
					errChans[i] <- nil
					return
				}
				if err != nil {
					errChans[i] <- err
					return
				}
				for _, title := range titles {
					title.SlugGroup = slug.Make(title.Name)
					title.Slug = title.SlugGroup
					if err := dao.Save(title); err != nil {
						errChans[i] <- err
						return
					}
				}
				errChans[i] <- nil
			})()
		}

		for _, errChan := range errChans {
			err := <-errChan
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db dbx.Builder) error {
		return nil
	})
}

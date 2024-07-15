package migrations

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func fetchAllReleases(dao *daos.Dao) (releases []*models.Release, err error) {
	err = models.ReleaseQuery(dao).All(&releases)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return
}

func init() {
	// will it ever be too big and I need to fix this for performance?
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		releases, err := fetchAllReleases(dao)
		if err != nil {
			return err
		}
		for _, release := range releases {
			if strings.Contains(release.Type, "(Digital)") {
				release.Type = strings.TrimSpace(strings.ReplaceAll(release.Type, "(Digital)", ""))
				release.Digital = true
				if err := dao.Save(release); err != nil {
					return err
				}
			}
		}
		return nil
	}, func(db dbx.Builder) error {
		return nil
	})
}

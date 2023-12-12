package migrations

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
	"tana.moe/momoka-lite/services"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		titleTableName := (&models.Title{}).TableName()
		duplicatedTitles := []*models.Title{}
		err := models.TitleQuery(dao).
			Select(
				fmt.Sprintf("%s.slugGroup", titleTableName),
				fmt.Sprintf("COUNT(%s.slug) AS count", titleTableName),
			).
			Having(
				dbx.NewExp("count > 1"),
			).
			GroupBy(fmt.Sprintf("%s.slugGroup", titleTableName)).
			All(&duplicatedTitles)
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		if err != nil {
			return err
		}
		for _, duplicatedSlug := range duplicatedTitles {
			if err := services.UpdateTitleSlug(dao, duplicatedSlug.SlugGroup); err != nil {
				return err
			}
		}
		return nil
	}, func(db dbx.Builder) error {
		return nil
	})
}

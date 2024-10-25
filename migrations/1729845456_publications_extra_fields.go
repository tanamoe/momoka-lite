package migrations

import (
	"regexp"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		publications := []*models.Publication{}
		if err := models.PublicationQuery(dao).All(&publications); err != nil {
			return err
		}
		for _, publication := range publications {
			// match last bracket and turn it into subtitle
			if strings.HasSuffix(publication.Name, ")") {
				reg := regexp.MustCompile(`(?s)\((.*)\)$`)
				match := reg.FindAllStringSubmatch(publication.Name, -1)
				publication.Subtitle = match[0][1]
				publication.Name = reg.ReplaceAllString(publication.Name, "")
			}

			// trim outdated reprint string
			if strings.Contains(publication.Subtitle, "Tái bản") {
				publication.Subtitle = ""
			}

			if err := dao.Save(publication); err != nil {
				return err
			}
		}

		return nil
	}, func(db dbx.Builder) error {
		return nil
	})
}

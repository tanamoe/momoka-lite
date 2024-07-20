package migrations

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"tana.moe/momoka-lite/models"
)

func findPublicationDefaultBook(dao *daos.Dao, publication *models.Publication) (*models.Book, error) {
	book := &models.Book{}
	err := models.BookQuery(dao).Where(&dbx.HashExp{
		"publication": publication.Id,
		"edition":     "",
	}).One(book)
	if err == nil {
		return book, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	err = models.BookQuery(dao).Where(&dbx.HashExp{
		"publication": publication.Id,
	}).One(book)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return book, err
}

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		publications := []*models.Publication{}
		if err := models.PublicationQuery(dao).All(&publications); err != nil {
			return err
		}
		for _, publication := range publications {
			defaultBook, err := findPublicationDefaultBook(dao, publication)
			if err != nil {
				return err
			}
			if defaultBook == nil {
				continue
			}
			publication.DefaultBookId = defaultBook.Id
			if err := dao.Save(publication); err != nil {
				return err
			}
		}
		return nil
	}, func(db dbx.Builder) error {
		return nil
	})
}

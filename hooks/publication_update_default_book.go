package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
)

func registerPublicationUpdateDefaultBook(
	app *pocketbase.PocketBase,
	context *models.AppContext,
) error {
	app.
		// TODO: convert models.Book into "books" collection
		OnModelAfterCreate("books").
		Add(
			func(e *core.ModelEvent) error {
				bookId := e.Model.GetId()
				book, err := models.FindBookById(app.Dao(), bookId)
				if err != nil {
					return err
				}
				if err := book.Expand(app.Dao(), models.ExpandMap{
					"publication": {},
				}); err != nil {
					return err
				}
				if book.Publication.DefaultBookId == "" {
					publication, err := models.FindPublicationById(app.Dao(), book.PublicationID)
					if err != nil {
						return err
					}
					publication.DefaultBookId = bookId
					if err := app.Dao().WithoutHooks().Save(publication); err != nil {
						return nil
					}
				}

				return nil
			},
		)
	return nil
}

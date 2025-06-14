package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"tana.moe/momoka-lite/models"
)

func registerPublicationUpdateDefaultBook(
	app *pocketbase.PocketBase,
) error {
	app.
		// TODO: convert models.Book into "books" collection
		OnModelAfterCreateSuccess("books").
		BindFunc(
			func(e *core.ModelEvent) error {
				bookId := e.Model.PK().(string)
				book, err := models.FindBookById(app.DB(), bookId)
				if err != nil {
					return err
				}
				if err := book.Expand(app.DB(), models.ExpandMap{
					"publication": {},
				}); err != nil {
					return err
				}
				if book.Publication.DefaultBookId == "" {
					publication, err := models.FindPublicationById(app.DB(), book.PublicationID)
					if err != nil {
						return err
					}
					publication.DefaultBookId = bookId
					if err := app.UnsafeWithoutHooks().DB().Model(publication).Update(); err != nil {
						return nil
					}
				}

				return e.Next()
			},
		)
	return nil
}

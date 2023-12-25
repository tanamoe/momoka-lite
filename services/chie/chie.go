package chie

import (
	"github.com/pocketbase/pocketbase"
	"tana.moe/momoka-lite/models"
)

type QueryKind string

const (
	AndQuery QueryKind = "and"
	OrQuery  QueryKind = "or"
)

type QueryGroup[T comparable] struct {
	Kind   QueryKind `json:"kind"`
	Values []T       `json:"values"`
}

func StartService(app *pocketbase.PocketBase, context *models.AppContext) error {
	context.Logger.Info("Indexing collections...")
	if err := startIndexing(app, context); err != nil {
		return err
	}
	context.Logger.Info("Collections indexing completed.")
	return nil
}

func startIndexing(app *pocketbase.PocketBase, context *models.AppContext) error {
	if err := indexTitleCollection(app, context); err != nil {
		return err
	}
	return nil
}

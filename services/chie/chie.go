package chie

import (
	"fmt"

	"github.com/blevesearch/bleve/v2/search/query"
	"github.com/pocketbase/pocketbase"
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

func queryGroupValues[T comparable](qg QueryGroup[T]) (vals []string) {
	for _, val := range qg.Values {
		vals = append(vals, fmt.Sprintf("%v", val))
	}
	return vals
}

func groupQuery[T comparable](qg QueryGroup[T], queries []query.Query) query.Query {
	switch qg.Kind {
	case AndQuery:
		return query.NewConjunctionQuery(queries)
	case OrQuery:
		return query.NewDisjunctionQuery(queries)
	}
	return query.NewMatchNoneQuery()
}

func groupMatchQuery[T comparable](qg QueryGroup[T], fields ...string) query.Query {
	var queries []query.Query
	for _, val := range queryGroupValues(qg) {
		if len(fields) > 0 {
			var cols []query.Query
			for _, field := range fields {
				q := query.NewMatchQuery(val)
				q.SetField(field)
				cols = append(cols, q)
			}
			queries = append(queries, query.NewDisjunctionQuery(cols))
			continue
		}
		q := query.NewMatchQuery(val)
		queries = append(queries, q)
	}
	return groupQuery(qg, queries)
}

func groupTermQuery[T comparable](qg QueryGroup[T], fields ...string) query.Query {
	var queries []query.Query
	for _, val := range queryGroupValues(qg) {
		if len(fields) > 0 {
			var cols []query.Query
			for _, field := range fields {
				q := query.NewTermQuery(val)
				q.SetField(field)
				cols = append(cols, q)
			}
			queries = append(queries, query.NewDisjunctionQuery(cols))
			continue
		}
		q := query.NewTermQuery(val)
		queries = append(queries, q)
	}
	return groupQuery(qg, queries)
}

func StartService(app *pocketbase.PocketBase) error {
	app.Logger().Info("Indexing collections...")
	if err := startIndexing(app); err != nil {
		return err
	}
	app.Logger().Info("Collections indexing completed.")
	app.Logger().Info("Starting chie signal gateways...")
	if err := startTitleSearchService(app); err != nil {
		return err
	}
	if err := startReleaseSearchService(app); err != nil {
		return err
	}
	app.Logger().Info("Chie signal gateways is now ready.")
	return nil
}

func startIndexing(app *pocketbase.PocketBase) error {
	if err := indexTitleCollection(app); err != nil {
		return err
	}
	if err := indexReleaseCollection(app); err != nil {
		return err
	}
	return nil
}

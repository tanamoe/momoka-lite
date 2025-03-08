package chie

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/query"
	"github.com/gosimple/slug"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"tana.moe/momoka-lite/models"
)

type ReleaseSearchRequest struct {
	Name      string             `json:"name"`
	Publisher QueryGroup[string] `json:"publisher"`
	Partner   QueryGroup[string] `json:"partner"`
	Status    string             `json:"status"`
	Limit     int                `json:"limit"`
	Offset    int                `json:"offset"`
	SortBy    string             `json:"sort"`
}

type ReleaseSearchResponse struct {
	TotalItems int      `json:"totalItems"`
	Items      []string `json:"items"`
	Error      error    `json:"-"`
}

type releaseIndex struct {
	Id             string   `json:"id" db:"id"`
	Name           string   `json:"name" db:"name"`
	Type           string   `json:"type" db:"type"`
	Digital        bool     `json:"digital" db:"digital"`
	Disambiguation string   `json:"disambiguation" db:"disambiguation"`
	Tags           []string `json:"tags" db:"-"`
	Publisher      string   `json:"publisher" db:"publisher"`
	Partner        string   `json:"partner" db:"partner"`
	Status         string   `json:"status" db:"status"`
	Created        string   `json:"created" db:"created"`
	Updated        string   `json:"updated" db:"updated"`
}

type releaseSearchSignal struct {
	Query  ReleaseSearchRequest
	Result chan ReleaseSearchResponse
}

type releaseUpdateSignal struct {
	Db      dbx.Builder
	Release *models.Release
	Err     chan error
}

var (
	releaseSearchChannel = make(chan releaseSearchSignal)
	releaseUpdateChannel = make(chan releaseUpdateSignal)
	releaseIndexMapping  bleve.Index
)

func indexReleaseCollection(app *pocketbase.PocketBase) error {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		return err
	}
	releaseIndexMapping = index

	pageSize := int64(100)
	for offset := int64(0); ; offset += pageSize {
		releases, err := fetchReleasesList(app.DB(), offset, pageSize)
		if err != nil {
			return err
		}
		for _, release := range releases {
			if err := releaseIndexMapping.Index(release.Id, release); err != nil {
				return err
			}
		}
		if int64(len(releases)) < pageSize {
			return nil
		}
	}
}

func startReleaseSearchService(app *pocketbase.PocketBase) error {
	go func() {
		for {
			select {
			case signal := <-releaseSearchChannel:
				count, items, err := searchForReleases(signal.Query)
				signal.Result <- ReleaseSearchResponse{
					TotalItems: count,
					Items:      items,
					Error:      err,
				}

			case signal := <-releaseUpdateChannel:
				signal.Err <- updateReleaseIndex(signal.Db, signal.Release)
			}
		}
	}()
	return nil
}

func SearchForReleases(query ReleaseSearchRequest) ReleaseSearchResponse {
	res := make(chan ReleaseSearchResponse)
	releaseSearchChannel <- releaseSearchSignal{
		Query:  query,
		Result: res,
	}
	return <-res
}

func UpdateReleaseIndex(db dbx.Builder, release *models.Release) error {
	err := make(chan error)
	releaseUpdateChannel <- releaseUpdateSignal{
		Db:      db,
		Release: release,
		Err:     err,
	}
	return <-err
}

func searchForReleases(req ReleaseSearchRequest) (int, []string, error) {
	var queries []query.Query
	if req.Name != "" {
		q := query.NewMatchQuery(normalizeReleaseName(req.Name))
		q.SetField("tags")
		queries = append(queries, q)
	}
	if len(req.Publisher.Values) > 0 {
		queries = append(queries, groupTermQuery[string](req.Publisher, "publisher"))
	}
	if len(req.Partner.Values) > 0 {
		queries = append(queries, groupTermQuery[string](req.Partner, "partner"))
	}
	if req.Status != "" {
		q := query.NewMatchQuery(req.Status)
		q.SetField("status")
		queries = append(queries, q)
	}
	searchQuery := query.NewConjunctionQuery(queries)
	searchReq := bleve.NewSearchRequest(searchQuery)
	if req.Limit <= 0 {
		req.Limit = 50
	}
	if req.Limit > 200 {
		req.Limit = 200
	}
	searchReq.Size = req.Limit
	searchReq.From = req.Offset
	switch req.SortBy {
	case "name", "-name", "created", "-created", "updated", "-updated":
		searchReq.SortBy([]string{req.SortBy})
	}
	result, err := releaseIndexMapping.Search(searchReq)
	if err != nil {
		return 0, nil, err
	}
	var hits []string
	for _, hit := range result.Hits {
		hits = append(hits, hit.ID)
	}
	return int(result.Total), hits, nil
}

func updateReleaseIndex(db dbx.Builder, release *models.Release) error {
	releaseIdx := &releaseIndex{}
	err := models.
		ReleaseQuery(db).
		Where(
			dbx.HashExp{
				"id": release.Id,
			},
		).
		One(releaseIdx)
	if errors.Is(err, sql.ErrNoRows) {
		_ = releaseIndexMapping.Delete(release.Id)
		return nil
	}
	if err != nil {
		return err
	}
	releases, err := normalizeReleasesList(db, []*releaseIndex{releaseIdx})
	if err != nil {
		return err
	}
	releaseIdx = releases[0]
	if err := releaseIndexMapping.Index(releaseIdx.Id, releaseIdx); err != nil {
		return err
	}
	return nil
}

func fetchReleasesList(db dbx.Builder, offset int64, limit int64) (releases []*releaseIndex, err error) {
	err = models.ReleaseQuery(db).Offset(offset).Limit(limit).All(&releases)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	releases, err = normalizeReleasesList(db, releases)
	return
}

func normalizeReleasesList(db dbx.Builder, releases []*releaseIndex) ([]*releaseIndex, error) {
	for _, release := range releases {
		if err := updateReleaseIndexTags(release); err != nil {
			return nil, err
		}
	}
	return releases, nil
}

func updateReleaseIndexTags(index *releaseIndex) error {
	index.Tags = []string{
		index.Name,
		normalizeReleaseName(index.Name),
		// tagging disambiguation
		index.Disambiguation,
		normalizeReleaseName(index.Disambiguation),
	}
	return nil
}

func normalizeReleaseName(name string) string {
	return strings.ReplaceAll(
		slug.Make(name),
		"-",
		" ",
	)
}

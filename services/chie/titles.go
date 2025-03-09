package chie

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/query"
	"github.com/gosimple/slug"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/types"
	"tana.moe/momoka-lite/models"
)

type TitleSearchRequest struct {
	Name        string             `json:"name"`
	Format      QueryGroup[string] `json:"format"`
	Demographic QueryGroup[string] `json:"demographic"`
	Genres      QueryGroup[string] `json:"genres"`
	Staffs      QueryGroup[string] `json:"staffs"`
	Limit       int                `json:"limit"`
	Offset      int                `json:"offset"`
	SortBy      string             `json:"sort"`
}

type TitleSearchResponse struct {
	TotalItems int      `json:"totalItems"`
	Items      []string `json:"items"`
	Error      error    `json:"-"`
}

type titleIndex struct {
	Id             string                  `json:"id" db:"id"`
	Name           string                  `json:"name" db:"name"`
	Tags           []string                `json:"tags" db:"-"`
	Format         string                  `json:"format" db:"format"`
	Demographic    string                  `json:"demographic" db:"demographic"`
	Genres         types.JSONArray[string] `json:"genres" db:"genres"`
	Staffs         []string                `json:"staffs" db:"staffs"`
	AdditionalName []string                `json:"additionalName" db:"additionalName"`
	Created        string                  `json:"created" db:"created"`
	Updated        string                  `json:"updated" db:"updated"`
}

type titleSearchSignal struct {
	Query  TitleSearchRequest
	Result chan TitleSearchResponse
}

type titleUpdateSignal struct {
	Db    dbx.Builder
	Title *models.Title
	Err   chan error
}

var (
	titleSearchChannel = make(chan titleSearchSignal)
	titleUpdateChannel = make(chan titleUpdateSignal)
	titleIndexMapping  bleve.Index
)

func indexTitleCollection(app *pocketbase.PocketBase) error {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		return err
	}
	titleIndexMapping = index

	db := app.DB()
	pageSize := int64(100)
	for offset := int64(0); ; offset += pageSize {
		titles, err := fetchTitlesList(db, offset, pageSize)
		if err != nil {
			return err
		}
		for _, title := range titles {
			if err := titleIndexMapping.Index(title.Id, title); err != nil {
				return err
			}
		}
		if int64(len(titles)) < pageSize {
			return nil
		}
	}
}

func startTitleSearchService(app *pocketbase.PocketBase) error {
	go func() {
		for {
			select {
			case signal := <-titleSearchChannel:
				count, items, err := searchForTitles(signal.Query)
				signal.Result <- TitleSearchResponse{
					TotalItems: count,
					Items:      items,
					Error:      err,
				}

			case signal := <-titleUpdateChannel:
				signal.Err <- updateTitleIndex(signal.Db, signal.Title)
			}
		}
	}()
	return nil
}

func SearchForTitles(query TitleSearchRequest) TitleSearchResponse {
	res := make(chan TitleSearchResponse)
	titleSearchChannel <- titleSearchSignal{
		Query:  query,
		Result: res,
	}
	return <-res
}

func UpdateTitleIndex(db dbx.Builder, title *models.Title) error {
	err := make(chan error)
	titleUpdateChannel <- titleUpdateSignal{
		Db:    db,
		Title: title,
		Err:   err,
	}
	return <-err
}

func searchForTitles(req TitleSearchRequest) (int, []string, error) {
	var queries []query.Query
	if req.Name != "" {
		q := query.NewMatchQuery(normalizeTitleName(req.Name))
		q.SetField("tags")
		queries = append(queries, q)
	}
	if len(req.Format.Values) > 0 {
		queries = append(queries, groupTermQuery[string](req.Format, "format"))
	}
	if len(req.Demographic.Values) > 0 {
		queries = append(queries, groupTermQuery[string](req.Demographic, "demographic"))
	}
	if len(req.Genres.Values) > 0 {
		queries = append(queries, groupTermQuery[string](req.Genres, "genres"))
	}
	if len(req.Staffs.Values) > 0 {
		staffsQuery := QueryGroup[string]{
			Kind: req.Staffs.Kind,
		}
		for _, staff := range req.Staffs.Values {
			staffsQuery.Values = append(staffsQuery.Values, normalizeTitleStaffName(staff))
		}
		queries = append(queries, groupMatchQuery[string](staffsQuery, "staffs"))
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
	result, err := titleIndexMapping.Search(searchReq)
	if err != nil {
		return 0, nil, err
	}
	var hits []string
	for _, hit := range result.Hits {
		hits = append(hits, hit.ID)
	}
	return int(result.Total), hits, nil
}

func updateTitleIndex(db dbx.Builder, title *models.Title) error {
	titleIdx := &titleIndex{}
	err := models.
		TitleQuery(db).
		Where(
			dbx.HashExp{
				"id": title.Id,
			},
		).
		One(titleIdx)
	if errors.Is(err, sql.ErrNoRows) {
		_ = titleIndexMapping.Delete(title.Id)
		return nil
	}
	if err != nil {
		return err
	}
	titles, err := normalizeTitlesList(db, []*titleIndex{titleIdx})
	if err != nil {
		return err
	}
	titleIdx = titles[0]
	if err := titleIndexMapping.Index(titleIdx.Id, titleIdx); err != nil {
		return err
	}
	return nil
}

func fetchTitlesList(db dbx.Builder, offset int64, limit int64) (titles []*titleIndex, err error) {
	err = models.TitleQuery(db).Offset(offset).Limit(limit).All(&titles)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	titles, err = normalizeTitlesList(db, titles)
	return
}

func normalizeTitlesList(db dbx.Builder, titles []*titleIndex) ([]*titleIndex, error) {
	if err := assignStaffsToTitles(db, titles); err != nil {
		return nil, err
	}
	if err := assignAdditionalNameToTitles(db, titles); err != nil {
		return nil, err
	}
	for _, title := range titles {
		if err := updateTitleIndexTags(title); err != nil {
			return nil, err
		}
	}
	return titles, nil
}

func assignStaffsToTitles(db dbx.Builder, titles []*titleIndex) error {
	titleIds := []any{}
	for _, title := range titles {
		titleIds = append(titleIds, title.Id)
	}

	type StaffAndTitle struct {
		TitleId   string `db:"title"`
		StaffName string `db:"staffName"`
	}
	staffAndTitleMap := []*StaffAndTitle{}
	worksTableName := (&models.Work{}).TableName()
	staffsTableName := (&models.Staff{}).TableName()
	err := models.WorkQuery(db).
		Select(
			fmt.Sprintf("%s.title AS title", worksTableName),
			fmt.Sprintf("%s.name AS staffName", staffsTableName),
		).
		RightJoin(
			staffsTableName,
			dbx.NewExp(fmt.Sprintf("%s.staff = %s.id", worksTableName, staffsTableName)),
		).
		GroupBy(
			fmt.Sprintf("%s.title", worksTableName),
			fmt.Sprintf("%s.name", staffsTableName),
		).
		Where(
			dbx.HashExp{
				fmt.Sprintf("%s.title", worksTableName): titleIds,
			},
		).
		All(&staffAndTitleMap)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}

	titleIdMap := map[string]int{}
	for index, title := range titles {
		titleIdMap[title.Id] = index
	}

	for _, item := range staffAndTitleMap {
		titles[titleIdMap[item.TitleId]].Staffs = append(
			titles[titleIdMap[item.TitleId]].Staffs,
			normalizeTitleStaffName(item.StaffName),
		)
	}
	return nil
}

func assignAdditionalNameToTitles(db dbx.Builder, titles []*titleIndex) error {
	titleIds := []any{}
	for _, title := range titles {
		titleIds = append(titleIds, title.Id)
	}

	additionalNameMap := []*models.AdditionalTitleName{}
	err := models.AdditionalTitleNameQuery(db).Select().All(&additionalNameMap)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}

	titleIdMap := map[string]int{}
	for index, title := range titles {
		titleIdMap[title.Id] = index
	}

	for _, item := range additionalNameMap {
		titles[titleIdMap[item.TitleId]].AdditionalName = append(
			titles[titleIdMap[item.TitleId]].AdditionalName,
			normalizeTitleName(item.Name),
		)
	}
	return nil
}

func normalizeTitleStaffName(name string) string {
	return strings.ReplaceAll(
		slug.Make(name),
		"-",
		" ",
	)
}

func updateTitleIndexTags(index *titleIndex) error {
	index.Tags = []string{
		index.Name,
		normalizeTitleName(index.Name),
	}
	index.Tags = append(index.Tags, index.AdditionalName...)
	index.Tags = append(index.Tags, normalizeAdditionalName(index.AdditionalName)...)
	return nil
}

func normalizeTitleName(name string) string {
	return strings.ReplaceAll(
		slug.Make(name),
		"-",
		" ",
	)
}

func normalizeAdditionalName(additionalName []string) []string {
	normalizedAdditionalName := []string{}
	for _, name := range additionalName {
		normalizedAdditionalName = append(
			normalizedAdditionalName,
			strings.ReplaceAll(
				slug.Make(name),
				"-",
				" ",
			),
		)
	}
	return normalizedAdditionalName
}

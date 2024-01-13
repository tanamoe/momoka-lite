package chie

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/gosimple/slug"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/daos"
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
}

type TitleSearchResponse struct {
	TotalItems int      `json:"totalItems"`
	Items      []string `json:"items"`
	Error      error    `json:"-"`
}

type titleIndex struct {
	Id          string                  `json:"id" db:"id"`
	Name        string                  `json:"name" db:"name"`
	Tags        []string                `json:"tags" db:"-"`
	Format      string                  `json:"format" db:"format"`
	Demographic string                  `json:"demographic" db:"demographic"`
	Genres      types.JsonArray[string] `json:"genres" db:"genres"`
	Staffs      []string                `json:"staffs" db:"staffs"`
}

type titleSearchSignal struct {
	Query  TitleSearchRequest
	Result chan TitleSearchResponse
}

type titleUpdateSignal struct {
	Dao   *daos.Dao
	Title *models.Title
	Err   chan error
}

var (
	titleSearchChannel = make(chan titleSearchSignal)
	titleUpdateChannel = make(chan titleUpdateSignal)
	titleIndexMapping  bleve.Index
)

func indexTitleCollection(app *pocketbase.PocketBase, context *models.AppContext) error {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		return err
	}
	titleIndexMapping = index

	dao := app.Dao()
	pageSize := int64(100)
	for offset := int64(0); ; offset += pageSize {
		titles, err := fetchTitlesList(dao, offset, pageSize)
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

func startTitleSearchService(app *pocketbase.PocketBase, context *models.AppContext) error {
	go func() {
		select {
		case signal := <-titleSearchChannel:
			count, items, err := searchForTitles(signal.Query)
			signal.Result <- TitleSearchResponse{
				TotalItems: count,
				Items:      items,
				Error:      err,
			}

		case signal := <-titleUpdateChannel:
			signal.Err <- updateTitleIndex(signal.Dao, signal.Title)
		}
	}()
	return nil
}

func searchForTitles(req TitleSearchRequest) (int, []string, error) {
	return 0, nil, nil
}

func updateTitleIndex(dao *daos.Dao, title *models.Title) error {
	titleIdx := &titleIndex{}
	err := models.
		TitleQuery(dao).
		Where(
			dbx.HashExp{
				"id": title.Id,
			},
		).
		One(titleIdx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}
	titles, err := normalizeTitlesList(dao, []*titleIndex{titleIdx})
	if err != nil {
		return err
	}
	titleIdx = titles[0]
	if err := titleIndexMapping.Index(titleIdx.Id, titleIdx); err != nil {
		return err
	}
	return nil
}

func fetchTitlesList(dao *daos.Dao, offset int64, limit int64) (titles []*titleIndex, err error) {
	err = models.TitleQuery(dao).Offset(offset).Limit(limit).All(&titles)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	titles, err = normalizeTitlesList(dao, titles)
	return
}

func normalizeTitlesList(dao *daos.Dao, titles []*titleIndex) ([]*titleIndex, error) {
	if err := assignStaffsToTitles(dao, titles); err != nil {
		return nil, err
	}
	for _, title := range titles {
		if err := updateTitleIndexTags(title); err != nil {
			return nil, err
		}
	}
	return titles, nil
}

func assignStaffsToTitles(dao *daos.Dao, titles []*titleIndex) error {
	titleIds := []any{}
	for _, title := range titles {
		titleIds = append(titleIds, title.Id)
	}

	type StaffAndTitle struct {
		TitleId   string `db:"title"`
		StaffId   string `db:"staffId"`
		StaffName string `db:"staffName"`
	}
	staffAndTitleMap := []*StaffAndTitle{}
	worksTableName := (&models.Work{}).TableName()
	staffsTableName := (&models.Staff{}).TableName()
	err := models.WorkQuery(dao).
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
				fmt.Sprintf("%s.id", worksTableName): titleIds,
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
	return nil
}

func normalizeTitleName(name string) string {
	return strings.ReplaceAll(
		slug.Make(name),
		"-",
		" ",
	)
}

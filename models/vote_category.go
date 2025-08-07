package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

type VoteBucketStrategy string

const (
	VoteSingleBucket VoteBucketStrategy = "single"
	VoteHourBucket   VoteBucketStrategy = "hour"
	VoteDayBucket    VoteBucketStrategy = "day"
	VoteWeekBucket   VoteBucketStrategy = "week"
	VoteMonthBucket  VoteBucketStrategy = "month"
	VoteYearBucket   VoteBucketStrategy = "year"
)

type VoteCategory struct {
	Id             string             `db:"id" json:"id"`
	Name           string             `db:"name" json:"name"`
	BucketStrategy VoteBucketStrategy `db:"bucketStrategy" json:"bucketStrategy"`
	CollectionId   string             `db:"collection" json:"collectionId"`
	Collection     *VoteCollection    `db:"-" json:"collection"`
	Start          types.DateTime     `db:"start" json:"start"`
	End            types.DateTime     `db:"end" json:"end"`
}

func (m *VoteCategory) TableName() string {
	return "voteCategories"
}

func VoteCategoryQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&VoteCategory{}).TableName())
}

func FindVoteCategoryById(db dbx.Builder, id string) (*VoteCategory, error) {
	category := &VoteCategory{}
	err := VoteCategoryQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(category)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (m *VoteCategory) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["collection"]; exist {
		collection, err := FindVoteCollectionById(db, m.CollectionId)
		if err != nil {
			return err
		}
		if collection != nil {
			if err := collection.Expand(db, e["collection"]); err != nil {
				return err
			}
			m.Collection = collection
		}
	}

	return nil
}

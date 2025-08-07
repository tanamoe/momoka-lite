package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

type VoteCollectionKind string

const (
	ManualVoteCollection  VoteCollectionKind = "manual"
	DynamicVoteCollection VoteCollectionKind = "dynamic"
)

type VoteCollection struct {
	Id     string             `db:"id" json:"id"`
	Kind   VoteCollectionKind `db:"kind" json:"kind"`
	Filter types.JSONRaw      `db:"filter" json:"filter"`
}

func (m *VoteCollection) TableName() string {
	return "voteCollections"
}

func VoteCollectionQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&VoteCollection{}).TableName())
}

func FindVoteCollectionById(db dbx.Builder, id string) (*VoteCollection, error) {
	collection := &VoteCollection{}
	err := VoteCollectionQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(collection)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (m *VoteCollection) Expand(db dbx.Builder, e ExpandMap) error {
	return nil
}

package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
)

type VoteCollectionItem struct {
	Id           string          `db:"id" json:"id"`
	CollectionId string          `db:"collection" json:"collectionId"`
	Collection   *VoteCollection `db:"-" json:"collection"`
	TitleId      string          `db:"title" json:"titleId"`
	Title        *Title          `db:"-" json:"title"`
}

func (m *VoteCollectionItem) TableName() string {
	return "voteCollectionItems"
}

func VoteCollectionItemQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&VoteCollectionItem{}).TableName())
}

func FindCollectionItemById(db dbx.Builder, id string) (*VoteCollectionItem, error) {
	item := &VoteCollectionItem{}
	err := VoteCollectionItemQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(item)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (m *VoteCollectionItem) Expand(db dbx.Builder, e ExpandMap) error {
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

	if _, exist := e["title"]; exist {
		title, err := FindTitleById(db, m.TitleId)
		if err != nil {
			return err
		}
		if title != nil {
			if err := title.Expand(db, e["title"]); err != nil {
				return err
			}
			m.Title = title
		}
	}

	return nil
}

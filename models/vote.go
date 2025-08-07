package models

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/dbx"
)

type Vote struct {
	Id         string        `db:"id" json:"id"`
	CategoryId string        `db:"category" json:"categoryId"`
	Category   *VoteCategory `db:"-" json:"category"`
	BucketId   string        `db:"bucketId" json:"bucketId"`
	UserId     string        `db:"user" json:"userId"`
	User       *User         `db:"-" json:"user"`
	TitleId    string        `db:"title" json:"titleId"`
	Title      *Title        `db:"-" json:"title"`
}

func (m *Vote) TableName() string {
	return "vote"
}

func VoteQuery(db dbx.Builder) *dbx.SelectQuery {
	return db.Select("*").From((&Vote{}).TableName())
}

func FindVoteById(db dbx.Builder, id string) (*Vote, error) {
	vote := &Vote{}
	err := VoteQuery(db).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(vote)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return vote, nil
}

func (m *Vote) Expand(db dbx.Builder, e ExpandMap) error {
	if e == nil {
		return nil
	}

	if _, exist := e["category"]; exist {
		category, err := FindVoteCategoryById(db, m.CategoryId)
		if err != nil {
			return err
		}
		if category != nil {
			if err := category.Expand(db, e["category"]); err != nil {
				return err
			}
			m.Category = category
		}
	}

	if _, exist := e["user"]; exist {
		user, err := FindUserById(db, m.UserId)
		if err != nil {
			return err
		}
		if user != nil {
			if err := user.Expand(db, e["user"]); err != nil {
				return err
			}
			m.User = user
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

package repo

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type RecordTimestamp struct {
	CreatedAt      time.Time
	LastModifiedAt bun.NullTime
	DeletedAt      bun.NullTime `bun:",soft_delete,nullzero"`
}

type CompactRecordTimestamp struct {
	CreatedAt      time.Time
	LastModifiedAt time.Time
	DeletedAt      bun.NullTime `bun:",soft_delete,nullzero"`
}

type CreateOnlyTimestamp struct {
	CreatedAt time.Time
}

var _ bun.BeforeAppendModelHook = (*RecordTimestamp)(nil)

func (recordTimestamp *RecordTimestamp) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		recordTimestamp.CreatedAt = time.Now()
		recordTimestamp.LastModifiedAt = bun.NullTime{Time: time.Now()}
	case *bun.UpdateQuery:
		recordTimestamp.LastModifiedAt = bun.NullTime{Time: time.Now()}
	}
	return nil
}

var _ bun.BeforeAppendModelHook = (*CompactRecordTimestamp)(nil)

func (recordTimestamp *CompactRecordTimestamp) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		currentTime := time.Now()
		recordTimestamp.CreatedAt = currentTime
		recordTimestamp.LastModifiedAt = currentTime
	case *bun.UpdateQuery:
		recordTimestamp.LastModifiedAt = time.Now()
	}
	return nil
}

var _ bun.BeforeAppendModelHook = (*CreateOnlyTimestamp)(nil)

func (recordTimestamp *CreateOnlyTimestamp) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		currentTime := time.Now()
		recordTimestamp.CreatedAt = currentTime
	}

	return nil
}

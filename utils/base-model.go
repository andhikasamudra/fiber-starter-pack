package utils

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type BaseModel struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	GUID      uuid.UUID `bun:"default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `bun:",soft_delete,nullzero"`
}

func (m *BaseModel) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	now := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = now
	case *bun.UpdateQuery:
		m.UpdatedAt = &now
	}
	return nil
}

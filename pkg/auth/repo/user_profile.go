package repo

import (
	"context"

	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	"github.com/uptrace/bun"
)

type UserProfile struct {
	ID          int `bun:",pk,autoincrement"`
	UserID      int
	Name        string
	PhoneNumber string
	CompanyName string
	RecordTimestamp

	User User `bun:"rel:belongs-to,join:user_id=id"`

	bun.BaseModel `bun:"table:user_profile,alias:up"`
}

func (r *Repo) CreateUserProfile(ctx context.Context, db postgres.Interface, data UserProfile) (*UserProfile, error) {
	r.Mtx.Lock()
	defer r.Mtx.Unlock()

	_, err := db.GetConnection().NewInsert().Model(&data).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

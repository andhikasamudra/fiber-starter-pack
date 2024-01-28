package repo

import (
	"context"

	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	"github.com/uptrace/bun"
)

type UserRole struct {
	UserID  int    `bun:",pk"`
	RoleKey string `bun:",pk"`

	User User `bun:"rel:belongs-to,join:user_id=id"`
	Role Role `bun:"rel:belongs-to,join:role_key=key"`

	bun.BaseModel `bun:"table:user_roles,alias:ur"`
}

func (r *Repo) CreateUserRole(ctx context.Context, db postgres.Interface, data UserRole) (*UserRole, error) {
	_, err := db.GetConnection().NewInsert().Model(&data).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

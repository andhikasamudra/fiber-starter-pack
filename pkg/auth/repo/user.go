package repo

import (
	"context"

	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type GetUserRequest struct {
	Email string
	GUID  uuid.UUID
}

// User Constructs your User model under entities.
type User struct {
	ID       int       `bun:",pk,autoincrement"`
	GUID     uuid.UUID `bun:",nullzero"`
	Email    string
	IsActive bool
	RecordTimestamp

	Profile *UserProfile `bun:"rel:has-one,join:id=user_id"`
	Roles   []UserRole   `bun:"rel:has-many,join:id=user_id"`

	bun.BaseModel `bun:"table:users,alias:u"`
}

func (r *Repo) CreateUser(ctx context.Context, db postgres.Interface, data User) (*User, error) {
	r.Mtx.Lock()
	defer r.Mtx.Unlock()

	_, err := db.GetConnection().NewInsert().Model(&data).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Repo) ReadUser(ctx context.Context, db postgres.Interface, request GetUserRequest) (*User, error) {
	var user User

	query := db.Conn().NewSelect().Model(&user).
		Relation("Profile").
		Relation("Roles").
		Relation("Documents")

	if request.Email != "" {
		query.Where("email = ?", request.Email)
	}

	if request.GUID != uuid.Nil {
		query.Where("guid = ?", request.GUID)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repo) UpdateUser(ctx context.Context, db postgres.Interface, data User, updatedFields []string) error {
	r.Mtx.Lock()
	defer r.Mtx.Unlock()

	_, err := db.GetConnection().NewUpdate().Model(&data).Column(updatedFields...).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

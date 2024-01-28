package repo

import (
	"context"

	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
)

type Interface interface {
	CreateUser(ctx context.Context, db postgres.Interface, book User) (*User, error)
	ReadUser(ctx context.Context, db postgres.Interface, request GetUserRequest) (*User, error)
	CreateUserProfile(ctx context.Context, db postgres.Interface, data UserProfile) (*UserProfile, error)
	CreateUserRole(ctx context.Context, db postgres.Interface, data UserRole) (*UserRole, error)
	UpdateUser(ctx context.Context, db postgres.Interface, data User, updatedFields []string) error
}

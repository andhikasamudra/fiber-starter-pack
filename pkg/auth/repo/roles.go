package repo

import (
	"github.com/uptrace/bun"
)

const (
	RoleAirline = "role.airline"
	RoleAgent   = "role.agent"
	RoleAdmin   = "role.admin"
)

type Role struct {
	Key string `bun:",pk"`

	bun.BaseModel `bun:"table:roles,alias:r"`
}

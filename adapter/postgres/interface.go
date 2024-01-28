package postgres

import (
	"github.com/uptrace/bun"
)

type Interface interface {
	HealthCheck() error
	Commit() error
	Conn() *bun.DB
	Rollback() error
	BeginTransaction() error
	GetConnection() bun.IDB
	Close()
}

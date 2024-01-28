package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Adapter struct {
	Db *bun.DB
	Tx *bun.Tx
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

// HealthCheck ...
func (a *Adapter) HealthCheck() error {
	_, err := a.Db.Exec("SELECT 1")
	if err != nil {
		fmt.Println("PostgreSQL is down")

		return err
	}

	return nil
}

func (a *Adapter) BeginTransaction() error {
	tx, err := a.Db.Begin()
	if err != nil {
		return err
	}

	a.Tx = &tx

	return nil
}

func (a *Adapter) Commit() error {
	if a.Tx == nil {
		return fmt.Errorf("transaction not ready")
	}

	err := a.Tx.Commit()
	if err != nil {
		return err
	}

	a.Tx = nil

	return nil
}

func (a *Adapter) Rollback() error {
	if a.Tx == nil {
		return fmt.Errorf("transaction not ready")
	}

	err := a.Tx.Rollback()
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) Conn() *bun.DB {
	return a.Db
}

func (a *Adapter) GetConnection() bun.IDB {
	if a.Tx != nil {
		return a.Tx
	}

	return a.Db
}

func (a *Adapter) Connect() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_DB_USER"),
		os.Getenv("POSTGRES_DB_PASS"),
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_DB_PORT"),
		os.Getenv("POSTGRES_DB_NAME"),
	)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	a.Db = bun.NewDB(sqldb, pgdialect.New())
}

func (a *Adapter) Close() {
	a.Db.Close()
}

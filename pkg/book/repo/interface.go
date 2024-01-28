package repo

import (
	"context"
	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	"github.com/andhikasamudra/fiber-starter-pack/utils"
	"github.com/google/uuid"
)

type Interface interface {
	CreateBook(ctx context.Context, dbAdapter postgres.Interface, book Book) (*Book, error)
	GetBook(ctx context.Context, dbAdapter postgres.Interface, guid uuid.UUID) (*Book, error)
	GetBooks(ctx context.Context, dbAdapter postgres.Interface, param GetBooksParam) ([]Book, *utils.TableListParams, error)
	UpdateBook(ctx context.Context, dbAdapter postgres.Interface, data Book, updatedColumn []string) error
	DeleteBook(ctx context.Context, dbAdapter postgres.Interface, data Book) error
}

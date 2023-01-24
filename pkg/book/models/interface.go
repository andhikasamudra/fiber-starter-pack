package models

import (
	"context"
)

type BookInterface interface {
	CreateBook(ctx context.Context, book Book) (*Book, error)
	ReadBook(ctx context.Context) ([]Book, error)
	UpdateBook(ctx context.Context, book Book, updatedColumn []string) error
	DeleteBook(ctx context.Context, bookID int) error
}

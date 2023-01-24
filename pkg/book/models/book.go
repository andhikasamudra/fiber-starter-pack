package models

import (
	"context"
	"fiber-starter-pack/utils"
	"github.com/uptrace/bun"
)

// Book Constructs your Book model under entities.
type Book struct {
	utils.BaseModel
	Title  string
	Author string
}

type Model struct {
	db *bun.DB
}

// NewRepo is the single instance repo that is being created.
func NewModel(db *bun.DB) *Model {
	return &Model{
		db: db,
	}
}

// CreateBook is a mongo repository that helps to create books
func (r *Model) CreateBook(ctx context.Context, book Book) (*Book, error) {
	_, err := r.db.NewInsert().Model(&book).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &book, nil

}

// ReadBook is a mongo repository that helps to fetch books
func (r *Model) ReadBook(ctx context.Context) ([]Book, error) {
	var books []Book

	err := r.db.NewSelect().Model(&books).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

// UpdateBook is a mongo repository that helps to update books
func (r *Model) UpdateBook(ctx context.Context, book Book, updatedColumn []string) error {
	_, err := r.db.NewUpdate().
		Model(&book).
		Column(updatedColumn...).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// DeleteBook is a mongo repository that helps to delete books
func (r *Model) DeleteBook(ctx context.Context, bookID int) error {
	_, err := r.db.NewDelete().
		Where("id = ?", bookID).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

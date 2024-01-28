package repo

import (
	"context"
	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	"github.com/google/uuid"

	"github.com/andhikasamudra/fiber-starter-pack/utils"
)

type GetBooksParam struct {
	*utils.TableListParams
}

type Book struct {
	utils.BaseModel
	Title  string
	Author string
}

func (r *Repo) CreateBook(ctx context.Context, dbAdapter postgres.Interface, book Book) (*Book, error) {
	_, err := dbAdapter.GetConnection().NewInsert().Model(&book).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &book, nil

}

func (r *Repo) GetBook(ctx context.Context, dbAdapter postgres.Interface, guid uuid.UUID) (*Book, error) {
	var book Book

	err := dbAdapter.GetConnection().
		NewSelect().
		Model(&book).
		Where("guid = ?", guid).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *Repo) GetBooks(ctx context.Context, dbAdapter postgres.Interface, param GetBooksParam) ([]Book, *utils.TableListParams, error) {
	var books []Book

	query := dbAdapter.GetConnection().NewSelect().Model(&books)

	if param.TableListParams != nil {
		resultQuery, paginateErr := utils.WithTableListParams(ctx, query, param.TableListParams)
		if paginateErr != nil {
			return nil, nil, paginateErr
		}

		query = resultQuery
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, nil, err
	}

	return books, param.TableListParams, nil
}

func (r *Repo) UpdateBook(ctx context.Context, dbAdapter postgres.Interface, data Book, updatedColumn []string) error {
	_, err := dbAdapter.GetConnection().NewUpdate().
		Model(&data).
		Column(updatedColumn...).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteBook(ctx context.Context, dbAdapter postgres.Interface, data Book) error {
	_, err := dbAdapter.GetConnection().NewDelete().Model(&data).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

package service

import (
	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/dto"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/repo"
	"github.com/gofiber/fiber/v2"
)

type Dependency struct {
	BookModel   repo.Interface
	DBAdapter   postgres.Interface
	RedisClient *redis.Client
	Logger      *zap.Logger
}

type BookService struct {
	BookModel   repo.Interface
	DBAdapter   postgres.Interface
	RedisClient *redis.Client
	Logger      *zap.Logger
}

func NewService(d Dependency) *BookService {
	return &BookService{
		BookModel:   d.BookModel,
		DBAdapter:   d.DBAdapter,
		RedisClient: d.RedisClient,
		Logger:      d.Logger,
	}
}

func (s *BookService) CreateBook(ctx *fiber.Ctx, request dto.CreateBookRequest) (*dto.GetBookResponse, error) {
	book := repo.Book{
		Title:  request.Title,
		Author: request.Author,
	}

	result, err := s.BookModel.CreateBook(ctx.Context(), s.DBAdapter, book)
	if err != nil {
		return nil, err
	}

	return &dto.GetBookResponse{
		GUID:   result.GUID,
		Title:  result.Title,
		Author: result.Author,
	}, nil
}

func (s *BookService) GetBook(ctx *fiber.Ctx, guid uuid.UUID) (*dto.GetBookResponse, error) {
	result, err := s.BookModel.GetBook(ctx.Context(), s.DBAdapter, guid)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &dto.GetBookResponse{
		GUID:   result.GUID,
		Title:  result.Title,
		Author: result.Author,
	}, nil
}

func (s *BookService) UpdateBook(ctx *fiber.Ctx, request dto.UpdateBookRequest) error {
	book, err := s.BookModel.GetBook(ctx.Context(), s.DBAdapter, request.GUID)
	if err != nil {
		return err
	}

	err = s.BookModel.UpdateBook(ctx.Context(), s.DBAdapter, *book, request.GetUpdatedColumns())
	if err != nil {
		return err
	}

	return nil
}
func (s *BookService) DeleteBook(ctx *fiber.Ctx, guid uuid.UUID) error {
	book, err := s.BookModel.GetBook(ctx.Context(), s.DBAdapter, guid)
	if err != nil {
		return err
	}

	err = s.BookModel.DeleteBook(ctx.Context(), s.DBAdapter, *book)
	if err != nil {
		return err
	}

	return nil
}

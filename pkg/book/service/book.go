package service

import (
	"context"

	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/dto"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/models"
	"github.com/gofiber/fiber/v2"
)

type Dependency struct {
	BookModel models.BookInterface
}

type BookService struct {
	BookModel models.BookInterface
}

func NewService(d Dependency) *BookService {
	return &BookService{
		BookModel: d.BookModel,
	}
}

func (s *BookService) CreateBook(ctx *fiber.Ctx, request dto.CreateBookRequest) (*models.Book, error) {
	book := models.Book{
		Title:  request.Title,
		Author: request.Author,
	}
	return s.BookModel.CreateBook(ctx.Context(), book)
}

func (s *BookService) ReadBook(ctx *fiber.Ctx) ([]dto.GetBookResponse, error) {
	result, err := s.BookModel.ReadBook(ctx.Context())
	if err != nil {
		return nil, err
	}

	return buildGetResponse(result), nil
}

func buildGetResponse(books []models.Book) []dto.GetBookResponse {
	var result []dto.GetBookResponse
	for _, i := range books {
		result = append(result, dto.GetBookResponse{
			GUID:   i.GUID,
			Title:  i.Title,
			Author: i.Author,
		})
	}

	return result
}

func (s *BookService) UpdateBook(ctx context.Context, book models.Book) error {
	return s.BookModel.UpdateBook(ctx, book, []string{})
}
func (s *BookService) DeleteBook(ctx context.Context, bookID int) error {
	return s.BookModel.DeleteBook(ctx, bookID)
}

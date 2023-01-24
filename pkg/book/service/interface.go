package service

import (
	"context"
	"fiber-starter-pack/dto"
	"fiber-starter-pack/pkg/book/models"
	"github.com/gofiber/fiber/v2"
)

type BookServiceInterface interface {
	CreateBook(ctx *fiber.Ctx, request dto.CreateBookRequest) (*models.Book, error)
	ReadBook(ctx *fiber.Ctx) ([]dto.GetBookResponse, error)
	UpdateBook(ctx context.Context, book models.Book) error
	DeleteBook(ctx context.Context, bookID int) error
}

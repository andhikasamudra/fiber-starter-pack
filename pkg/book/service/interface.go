package service

import (
	"github.com/google/uuid"

	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/dto"
	"github.com/gofiber/fiber/v2"
)

type Interface interface {
	CreateBook(ctx *fiber.Ctx, request dto.CreateBookRequest) (*dto.GetBookResponse, error)
	GetBook(ctx *fiber.Ctx, guid uuid.UUID) (*dto.GetBookResponse, error)
	UpdateBook(ctx *fiber.Ctx, request dto.UpdateBookRequest) error
	DeleteBook(ctx *fiber.Ctx, guid uuid.UUID) error
}

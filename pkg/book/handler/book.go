package handler

import (
	internalError "github.com/andhikasamudra/fiber-starter-pack/internal/error"
	"net/http"

	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateBook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.CreateBookRequest

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		err = h.Validator.Struct(request)
		if err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		result, err := h.BookService.CreateBook(c, request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusCreated).JSON(h.Resp.SetOk(result))
	}
}

package handler

import (
	"net/http"

	"github.com/andhikasamudra/fiber-starter-pack/internal"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/dto"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/service"
	"github.com/gofiber/fiber/v2"
)

type Dependency struct {
	BookService service.BookServiceInterface
}

type Handler struct {
	BookService service.BookServiceInterface
}

func NewHandler(d Dependency) *Handler {
	return &Handler{
		BookService: d.BookService,
	}
}

// AddBook is handler/controller which creates Books in the BookShop
func (h *Handler) AddBook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.CreateBookRequest

		err := c.BodyParser(&request)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(internal.BuildResponse(err))
		}

		result, err := h.BookService.CreateBook(c, request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildResponse(err))
		}

		return c.JSON(internal.BuildResponse(result))
	}
}

// UpdateBook is handler/controller which updates data of Books in the BookShop
//
//	func UpdateBook(service book.Service) fiber.Handler {
//		return func(c *fiber.Ctx) error {
//			var requestBody entities.Book
//			err := c.BodyParser(&requestBody)
//			if err != nil {
//				c.Status(http.StatusBadRequest)
//				return c.JSON(presenter.BookErrorResponse(err))
//			}
//			result, err := service.UpdateBook(&requestBody)
//			if err != nil {
//				c.Status(http.StatusInternalServerError)
//				return c.JSON(presenter.BookErrorResponse(err))
//			}
//			return c.JSON(presenter.BookSuccessResponse(result))
//		}
//	}
//
// // RemoveBook is handler/controller which removes Books from the BookShop
//
//	func RemoveBook(service book.Service) fiber.Handler {
//		return func(c *fiber.Ctx) error {
//			var requestBody entities.DeleteRequest
//			err := c.BodyParser(&requestBody)
//			if err != nil {
//				c.Status(http.StatusBadRequest)
//				return c.JSON(presenter.BookErrorResponse(err))
//			}
//			bookID := requestBody.ID
//			err = service.RemoveBook(bookID)
//			if err != nil {
//				c.Status(http.StatusInternalServerError)
//				return c.JSON(presenter.BookErrorResponse(err))
//			}
//			return c.JSON(&fiber.Map{
//				"status": true,
//				"data":   "updated successfully",
//				"err":    nil,
//			})
//		}
//	}
//
// // GetBooks is handler/controller which lists all Books from the BookShop
func (h *Handler) GetBooks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := h.BookService.ReadBook(c)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(internal.BuildResponse(err))
		}

		return c.JSON(internal.BuildResponse(result))
	}
}

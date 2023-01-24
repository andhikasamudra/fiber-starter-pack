package book

import (
	"fiber-starter-pack/pkg/book/handler"
	"fiber-starter-pack/pkg/book/models"
	"fiber-starter-pack/pkg/book/service"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func InitRoute(r fiber.Router, db *bun.DB) {
	m := models.NewModel(db)
	s := service.NewService(service.ServiceDependency{
		BookModel: m,
	})
	h := handler.NewHandler(handler.HandlerDependency{
		BookService: s,
	})

	api := r.Group("/book")
	api.Post("/", h.AddBook())
	api.Get("/", h.GetBooks())
}

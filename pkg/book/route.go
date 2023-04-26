package book

import (
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/handler"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/models"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/service"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func InitRoute(r fiber.Router, db *bun.DB) {
	m := models.NewModel(db)
	s := service.NewService(service.Dependency{
		BookModel: m,
	})
	h := handler.NewHandler(handler.Dependency{
		BookService: s,
	})

	api := r.Group("/book")
	api.Post("/", h.AddBook())
	api.Get("/", h.GetBooks())
}

package main

import (
	"fiber-starter-pack/config"
	"fiber-starter-pack/pkg/book"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
)

func main() {
	db := config.GetConnection()
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Testis"))
	})
	api := app.Group("/api")
	book.InitRoute(api, db)
	log.Fatal(app.Listen(":8080"))
}

package book

import (
	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	internalLogger "github.com/andhikasamudra/fiber-starter-pack/internal/logger"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/handler"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/repo"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/service"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func InitRoute(r fiber.Router, db postgres.Interface, rc *redis.Client) {
	logger := internalLogger.CreateLogger()
	m := repo.NewRepo()
	s := service.NewService(service.Dependency{
		BookModel:   m,
		DBAdapter:   db,
		RedisClient: rc,
		Logger:      logger,
	})
	h := handler.NewHandler(handler.Dependency{
		BookService: s,
	})

	api := r.Group("/book")
	api.Post("/", h.CreateBook())
}

package auth

import (
	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	internalLogger "github.com/andhikasamudra/fiber-starter-pack/internal/logger"
	"github.com/andhikasamudra/fiber-starter-pack/internal/middleware"
	"github.com/andhikasamudra/fiber-starter-pack/internal/validator"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/handler"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/repo"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/service"
	"github.com/andhikasamudra/fiber-starter-pack/provider/mail"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func InitRoute(r fiber.Router, db postgres.Interface, rc *redis.Client) {
	v := validator.NewValidator()
	m := repo.NewRepo()
	logger := internalLogger.CreateLogger()
	mailProvider := mail.NewProvider()

	s := service.NewService(service.Dependency{
		AuthModel:    m,
		RedisClient:  rc,
		MailProvider: mailProvider,
		Logger:       logger,
		DbAdapter:    db,
	})
	h := handler.NewHandler(handler.Dependency{
		AuthService: s,
		Validator:   v,
	})

	api := r.Group("/auth")
	api.Post("/register", h.RegisterUser())
	api.Get("/activate", h.ActivateUserRequest())
	api.Post("/otp", h.OTPRequest())
	api.Post("/otp/verify", h.LoginWithOTP())

	api.Get("/profile", middleware.Protected(), h.GetProfile())
}

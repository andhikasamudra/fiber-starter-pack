package auth

import (
	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	internalLogger "github.com/andhikasamudra/fiber-starter-pack/internal/logger"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/handler"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/repo"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/service"
)

func InitRPCServerRoute(db postgres.Interface) *handler.AuthRPCServer {
	m := repo.NewRepo()
	logger := internalLogger.CreateLogger()
	s := service.NewService(service.Dependency{
		AuthModel: m,
		Logger:    logger,
		DbAdapter: db,
	})
	server := handler.NewAuthRPCServer(handler.AuthRPCServerDependency{
		AuthService: s,
		Logger:      logger,
	})

	return server
}

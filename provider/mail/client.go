package mail

import (
	internalLogger "github.com/andhikasamudra/fiber-starter-pack/internal/logger"
	"go.uber.org/zap"
)

type Provider struct {
	Logger *zap.Logger
}

func NewProvider() *Provider {
	logger := internalLogger.CreateLogger()

	return &Provider{
		Logger: logger,
	}
}

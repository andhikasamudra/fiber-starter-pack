package handler

import (
	"context"
	"fmt"
	fiberstarterpack "github.com/andhikasamudra/fiber-starter-pack/protobuf"

	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/service"
	cargomarketbid "github.com/andhikasamudra/fiber-starter-pack/protobuf"
	"go.uber.org/zap"
)

type AuthRPCServer struct {
	AuthService service.Interface
	Logger      *zap.Logger
	fiberstarterpack.UnimplementedAuthServiceServer
}

type AuthRPCServerDependency struct {
	AuthService service.Interface
	Logger      *zap.Logger
}

func NewAuthRPCServer(d AuthRPCServerDependency) *AuthRPCServer {
	return &AuthRPCServer{
		AuthService: d.AuthService,
		Logger:      d.Logger,
	}
}

func (s *AuthRPCServer) GetUserProfile(_ context.Context, req *cargomarketbid.GetUserProfileRequest) (*cargomarketbid.GetUserProfileResponse, error) {
	resp, err := s.AuthService.GetProfileRPC(req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("failed to validate token : %s", err))
		return nil, err
	}

	return resp, nil
}

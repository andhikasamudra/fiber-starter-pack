package service

import (
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/dto"
	cargomarketbid "github.com/andhikasamudra/fiber-starter-pack/protobuf"
	"github.com/gofiber/fiber/v2"
)

type Interface interface {
	RegisterUser(ctx *fiber.Ctx, request dto.CreateUserRequest) error
	OTPRequest(ctx *fiber.Ctx, request dto.LoginRequest) error
	LoginWithOTPRequest(ctx *fiber.Ctx, request dto.LoginWithOTPRequest) (*dto.LoginResponse, error)
	ActivateUser(ctx *fiber.Ctx, request dto.ActivateUserRequest) error
	GetProfile(ctx *fiber.Ctx) (*dto.GetProfileResponse, error)

	// RPC
	GetProfileRPC(request *cargomarketbid.GetUserProfileRequest) (*cargomarketbid.GetUserProfileResponse, error)
}

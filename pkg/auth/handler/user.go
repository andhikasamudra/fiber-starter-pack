package handler

import (
	"fmt"
	"net/http"

	internalError "github.com/andhikasamudra/fiber-starter-pack/internal/error"
	"github.com/andhikasamudra/fiber-starter-pack/internal/handler"
	"github.com/go-playground/validator/v10"

	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/dto"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/service"
	"github.com/gofiber/fiber/v2"
)

type Dependency struct {
	AuthService service.Interface
	Validator   *validator.Validate
	Resp        handler.ResponseInterface
}

type Handler struct {
	AuthService service.Interface
	Validator   *validator.Validate
	Resp        handler.ResponseInterface
}

func NewHandler(d Dependency) *Handler {
	resp := handler.NewResponse()
	return &Handler{
		AuthService: d.AuthService,
		Validator:   d.Validator,
		Resp:        resp,
	}
}

func (h *Handler) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.CreateUserRequest

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		err = h.Validator.Struct(request)
		if err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		err = h.AuthService.RegisterUser(c, request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusCreated).JSON(h.Resp.SetOk(nil))
	}
}

func (h *Handler) OTPRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.LoginRequest

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		err = h.Validator.Struct(request)
		if err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		err = h.AuthService.OTPRequest(c, request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(h.Resp.SetOk(nil))
	}
}

func (h *Handler) ActivateUserRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Query("code")
		if code == "" {
			errorMsg := fmt.Errorf("code is required")
			return c.Status(http.StatusBadRequest).
				JSON(h.Resp.SetError(errorMsg, internalError.RequestError, errorMsg.Error()))
		}

		var request dto.ActivateUserRequest
		request.Code = code

		err := h.AuthService.ActivateUser(c, request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(h.Resp.SetOk(nil))
	}
}

func (h *Handler) LoginWithOTP() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.LoginWithOTPRequest

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		err = h.Validator.Struct(request)
		if err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		result, err := h.AuthService.LoginWithOTPRequest(c, request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(h.Resp.SetOk(result))
	}
}

func (h *Handler) GetProfile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := h.AuthService.GetProfile(c)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(h.Resp.SetOk(result))
	}
}

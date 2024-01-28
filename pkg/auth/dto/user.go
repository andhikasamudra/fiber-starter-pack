package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Role        string `json:"role" validate:"required"`
}

type ActivateUserRequest struct {
	Code string `json:"code"`
}

type ConfirmRegisterRequest struct {
	Email   string `json:"email" validate:"required"`
	OTPCode int    `json:"otp_code" validate:"required"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required"`
}

type LoginWithOTPRequest struct {
	Email   string `json:"email" validate:"required,email"`
	OTPCode string `json:"otp_code" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type GetProfileResponse struct {
	Email       string    `json:"email"`
	GUID        uuid.UUID `json:"guid"`
	PICName     string    `json:"pic_name"`
	CompanyName string    `json:"company_name"`
	PhoneNumber string    `json:"phone_number"`
	Roles       []string  `json:"roles"`
}

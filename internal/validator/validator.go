package validator

import (
	"github.com/go-playground/validator/v10"
)

func Register(v *validator.Validate) {
	err := v.RegisterValidation("enum", ValidateEnum)
	if err != nil {
		return
	}
}

func NewValidator() *validator.Validate {
	v := validator.New()
	Register(v)
	return v
}

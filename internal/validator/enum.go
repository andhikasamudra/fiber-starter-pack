package validator

import "github.com/go-playground/validator/v10"

type EnumValid interface {
	Valid() bool
}

func ValidateEnum(fl validator.FieldLevel) bool {
	if enum, ok := fl.Field().Interface().(EnumValid); ok {
		return enum.Valid()
	}
	return false
}

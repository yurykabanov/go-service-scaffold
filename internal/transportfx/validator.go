package transportfx

import (
	"github.com/go-playground/validator/v10"
)

type echoValidatorWrapper struct {
	validator *validator.Validate
}

func (v *echoValidatorWrapper) Validate(value interface{}) error {
	return v.validator.Struct(value)
}

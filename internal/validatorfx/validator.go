package validatorfx

import (
	"github.com/go-playground/validator/v10"
)

func ValidatorProvider() *validator.Validate {
	v := validator.New()

	// configure validator

	return v
}

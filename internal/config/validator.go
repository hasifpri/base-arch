package internalconfig

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator(env *Env) *validator.Validate {
	return validator.New()
}

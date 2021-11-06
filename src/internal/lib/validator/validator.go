package validator

import (
	"errors"
	"strings"

	validatorPkg "github.com/go-playground/validator/v10"
)

// Validator wraps the go playground validator for the echo framework interface.
type Validator struct {
	validator *validatorPkg.Validate
}

// NewValidator creates a new validator.
func NewValidator() *Validator {
	return &Validator{validator: validatorPkg.New()}
}

// Validate implements the echo framework validator interface.
func (val *Validator) Validate(i interface{}) error {
	err := val.validator.Struct(i)
	if err == nil {
		return nil
	}

	err = errors.New(strings.ReplaceAll(err.Error(), "\n", ", "))

	return err
}

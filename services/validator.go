package services

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Validator struct
type Validator struct {
	validator *validator.Validate
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}

// Validate validates the structs
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func ValidateRequestBody(c *fiber.Ctx, dto interface{}, out interface{}) error {
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	v := NewValidator()
	if err := v.Validate(dto); err != nil {
		return errors.New(err.Error())
	}

	if err := c.BodyParser(&out); err != nil {
		return err
	}

	return nil
}

package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ReadRequestFromBody(ctx *fiber.Ctx, dto interface{}) error {
	if err := ctx.BodyParser(dto); err != nil {
		return err
	}

	return validate.Struct(dto)
}
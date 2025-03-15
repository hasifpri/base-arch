package internalconfig

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiber(env *Env) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      env.AppName,
		ErrorHandler: NewErrorHandler(),
		Prefork:      env.PreFork,
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		return ctx.Status(code).JSON(fiber.Map{
			"code":    code,
			"message": err.Error(),
		})
	}
}

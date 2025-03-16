package helpergenerator

import (
	"context"
	"github.com/gofiber/fiber/v2"
	coreenum "panel-ektensi/core/enum"
)

func DefaultContextGenerator(fiberCtx *fiber.Ctx) context.Context {

	ctx := fiberCtx.UserContext()

	ctx = context.WithValue(fiberCtx.UserContext(), string(coreenum.CTXEnumIDUserID), fiberCtx.Locals(string(coreenum.CTXEnumIDUserID)))
	ctx = context.WithValue(fiberCtx.UserContext(), string(coreenum.CTXEnumIDUserEmail), fiberCtx.Locals(string(coreenum.CTXEnumIDUserEmail)))
	ctx = context.WithValue(fiberCtx.UserContext(), string(coreenum.CTXEnumIDUserName), fiberCtx.Locals(string(coreenum.CTXEnumIDUserName)))

	return ctx
}

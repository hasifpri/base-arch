package route

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type RouteConfig struct {
	App *fiber.App

	// ADMIN ROUTE
	// END ADMIN ROUTE
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAdminRoute()
	c.SetupCustomerRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/health", func(c *fiber.Ctx) error {
		tr := otel.Tracer("panel-ektensi/http")
		ctx := c.UserContext()

		ctx, span := tr.Start(ctx, "HealthHandler")
		defer span.End()

		span.SetAttributes(
			attribute.String("http.route", "/health"),
			attribute.String("app.component", "health"),
		)

		// error
		if false {
			err := errors.New("simulated error")
			span.SetStatus(codes.Error, err.Error())
			return err
		}

		return c.SendString("OK")
	})
}

func (c *RouteConfig) SetupAdminRoute() {

	// Admin Route
}

func (c *RouteConfig) SetupCustomerRoute() {

}

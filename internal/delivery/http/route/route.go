package internaldeliveryhttproute

import (
	"github.com/gofiber/fiber/v2"
	internaldeliveryhttpadmin "panel-ektensi/internal/delivery/http/admin"
)

type RouteConfig struct {
	App *fiber.App

	// ADMIN ROUTE
	AdminHandlerAdmin *internaldeliveryhttpadmin.AdminHandler
	// END ADMIN ROUTE
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAdminRoute()
	c.SetupCustomerRoute()
}

func (c *RouteConfig) SetupGuestRoute() {

}

func (c *RouteConfig) SetupAdminRoute() {

	adminRoute := c.App.Group("/admin")

	// Admin Route
	adminRoute.Post("/", c.AdminHandlerAdmin.Create)
}

func (c *RouteConfig) SetupCustomerRoute() {

}

package route

import (
	"github.com/gofiber/fiber/v2"
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

}

func (c *RouteConfig) SetupAdminRoute() {

	// Admin Route
}

func (c *RouteConfig) SetupCustomerRoute() {

}

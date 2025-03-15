package internaldeliveryhttproute

import "github.com/gofiber/fiber/v2"

type RouteConfig struct {
	App *fiber.App
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAdminRoute()
	c.SetupCustomerRoute()
}

func (c *RouteConfig) SetupGuestRoute() {

}

func (c *RouteConfig) SetupAdminRoute() {

}

func (c *RouteConfig) SetupCustomerRoute() {

}

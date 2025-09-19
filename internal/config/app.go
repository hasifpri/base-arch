package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log/slog"
	internaldeliveryhttproute "panel-ektensi/internal/delivery/http/route"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *slog.Logger
	Validate *validator.Validate
	Viper    *viper.Viper
	Env      *Env
}

func Bootstrap(config *BootstrapConfig) {

	// setup repository

	// setup usecase

	// setup controller

	routeConfig := internaldeliveryhttproute.RouteConfig{
		App: config.App,
	}

	routeConfig.Setup()
}

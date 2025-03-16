package internalconfig

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log/slog"
	internaldeliveryhttpadmin "panel-ektensi/internal/delivery/http/admin"
	internaldeliveryhttproute "panel-ektensi/internal/delivery/http/route"
	internalrepository "panel-ektensi/internal/repository"
	internalusecaseadmin "panel-ektensi/internal/usecase/admin"
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
	adminRepository := internalrepository.NewAdminRepository(config.Log)

	// setup usecase ADMIN
	adminUseCase := internalusecaseadmin.NewAdminUseCase(config.DB, config.Log, config.Validate, adminRepository)

	// setup controller ADMIN
	adminHandler := internaldeliveryhttpadmin.NewAdminHandler(adminUseCase, config.Log)

	routeConfig := internaldeliveryhttproute.RouteConfig{
		App:               config.App,
		AdminHandlerAdmin: adminHandler,
	}

	routeConfig.Setup()
}

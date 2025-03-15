package main

import (
	"fmt"
	"os"
	helpermigration "panel-ektensi/helper/migration"
	internalconfig "panel-ektensi/internal/config"
)

func main() {
	viperConfig := internalconfig.NewViper()
	envConfig := internalconfig.NewEnv(viperConfig)
	logConfig := internalconfig.NewSlog(envConfig)
	dbConfig := internalconfig.NewDatabase(envConfig, logConfig)

	validateConfig := internalconfig.NewValidator(envConfig)
	fiberConfig := internalconfig.NewFiber(envConfig)

	// Migrate DB
	if envConfig.DBMigrate {
		helpermigration.AutoMigrate(dbConfig)
	}

	// Init Bootstrap Server
	internalconfig.Bootstrap(&internalconfig.BootstrapConfig{
		DB:       dbConfig,
		App:      fiberConfig,
		Log:      logConfig,
		Validate: validateConfig,
		Viper:    viperConfig,
		Env:      envConfig,
	})

	// Running App
	err := fiberConfig.Listen(fmt.Sprintf(":%d", envConfig.WebPort))
	if err != nil {
		logConfig.Error("Error starting server", err)
		os.Exit(1)
	}
}

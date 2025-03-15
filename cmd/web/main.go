package main

import (
	"fmt"
	"os"
	internalconfig "panel-ektensi/internal/config"
)

func main() {
	viperConfig := internalconfig.NewViper()
	envConfig := internalconfig.NewEnv(viperConfig)
	logConfig := internalconfig.NewSlog(envConfig)
	dbConfig := internalconfig.NewDatabase(envConfig, logConfig)

	validateConfig := internalconfig.NewValidator(envConfig)
	fiberConfig := internalconfig.NewFiber(envConfig)

	internalconfig.Bootstrap(&internalconfig.BootstrapConfig{
		DB:       dbConfig,
		App:      fiberConfig,
		Log:      logConfig,
		Validate: validateConfig,
		Viper:    viperConfig,
		Env:      envConfig,
	})

	err := fiberConfig.Listen(fmt.Sprintf(":%d", envConfig.WebPort))
	if err != nil {
		logConfig.Error("Error starting server", err)
		os.Exit(1)
	}
}

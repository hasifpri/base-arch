package config

import (
	"github.com/spf13/viper"
	"time"
)

type Env struct {
	AppName  string
	PreFork  bool
	LogLevel string
	WebPort  int

	DBMigrate  bool
	DBUser     string
	DBPass     string
	DBHost     string
	DBPort     string
	DBName     string
	DBSslMode  string
	DBIdleConn int
	DBMaxConn  int
	DBMaxLife  time.Duration

	OTelEndpoint string
}

func NewEnv(viper *viper.Viper) *Env {
	return &Env{
		AppName:  viper.GetString("APP_NAME"),
		PreFork:  viper.GetBool("PRE_FORK"),
		LogLevel: viper.GetString("LOG_LEVEL"),
		WebPort:  viper.GetInt("WEB_PORT"),

		DBMigrate:  viper.GetBool("DB_MIGRATE"),
		DBUser:     viper.GetString("DB_USER"),
		DBPass:     viper.GetString("DB_PASS"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBName:     viper.GetString("DB_NAME"),
		DBSslMode:  viper.GetString("DB_SSL_MODE"),
		DBIdleConn: viper.GetInt("DB_IDLE_CONN"),
		DBMaxConn:  viper.GetInt("DB_MAX_CONN"),
		DBMaxLife:  viper.GetDuration("DB_MAX_LIFE"),

		OTelEndpoint: viper.GetString("OTEL_ENDPOINT"),
	}
}

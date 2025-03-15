package internalconfig

import (
	"log/slog"
	"os"
)

func NewSlog(env *Env) *slog.Logger {

	logLevelInt := GetLevelLog(env.LogLevel)

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevelInt,
	}))

	return log
}

func GetLevelLog(logLevelStr string) slog.Level {
	var logLevelInt slog.Level
	if logLevelStr == "Debug" {
		logLevelInt = slog.LevelDebug
	} else if logLevelStr == "Info" {
		logLevelInt = slog.LevelInfo
	} else if logLevelStr == "Warn" {
		logLevelInt = slog.LevelWarn
	} else {
		logLevelInt = slog.LevelError
	}

	return logLevelInt
}

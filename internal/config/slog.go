package config

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

func NewSlog(env *Env) *slog.Logger {

	logLevelInt := GetLevelLog(env.LogLevel)

	// Type Log
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Jika gagal membuka file, fallback ke os.Stdout
		logFile = os.Stdout
	}
	writer := io.MultiWriter(os.Stdout, logFile)

	log := slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: logLevelInt,
	}))

	return log
}

func GetLevelLog(logLevelStr string) slog.Level {

	logLevelStr = strings.ToLower(logLevelStr)
	
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

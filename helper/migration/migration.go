package helpermigration

import (
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func AutoMigrate(DB *gorm.DB) {

	err := DB.AutoMigrate()
	if err != nil {
		slog.Error("Error migrate database", err)
		os.Exit(1)
	}
}

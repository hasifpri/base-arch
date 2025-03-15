package helpermigration

import (
	"gorm.io/gorm"
	"log/slog"
	"os"
	internalentity "panel-ektensi/internal/entity"
)

func AutoMigrate(DB *gorm.DB) {

	err := DB.AutoMigrate(
		&internalentity.Admin{},
	)
	if err != nil {
		slog.Error("Error migrate database", err)
		os.Exit(1)
	}
}

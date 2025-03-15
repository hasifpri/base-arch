package internalrepository

import (
	"log/slog"
	internalentity "panel-ektensi/internal/entity"
)

type AdminRepository struct {
	Repository[internalentity.Admin]
	Log *slog.Logger
}

func NewAdminRepository(log *slog.Logger) *AdminRepository {
	return &AdminRepository{
		Log: log,
	}
}

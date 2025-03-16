package internalusecaseadmin

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log/slog"
	"panel-ektensi/helper"
	helperexception "panel-ektensi/helper/exception"
	internalentity "panel-ektensi/internal/entity"
	internalmodeladminrequest "panel-ektensi/internal/model/admin/request"
	internalmodeladminresponse "panel-ektensi/internal/model/admin/response"
	internalrepository "panel-ektensi/internal/repository"
)

type AdminUseCase struct {
	DB              *gorm.DB
	Log             *slog.Logger
	Validate        *validator.Validate
	AdminRepository *internalrepository.AdminRepository
}

func NewAdminUseCase(
	db *gorm.DB,
	log *slog.Logger,
	validate *validator.Validate,
	adminRepository *internalrepository.AdminRepository,
) *AdminUseCase {

	return &AdminUseCase{
		DB:              db,
		Log:             log,
		Validate:        validate,
		AdminRepository: adminRepository,
	}

}

func (usecase *AdminUseCase) Create(ctx context.Context, request *internalmodeladminrequest.CreateAdminInfo) (response internalmodeladminresponse.CreateAdminResponse, exc *helperexception.Exception) {

	// Init Transaction
	tx := usecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	password := helper.HashPassword(request.Password)
	request.Password = password

	// Convert Entity
	entity := internalentity.ConvertModelToEntitiesAdmin(request)

	// Exec DB
	err := usecase.AdminRepository.Create(tx, entity)
	if err != nil {
		usecase.Log.Error("usecase.Create()", "AdminRepository.Create()", "Error", err)
		exc = helperexception.Internal("failed insert admin", err)
		return
	}

	// Commit Transaction
	if err = tx.Commit().Error; err != nil {
		usecase.Log.Error("usecase.Create()", "tx.Commit()", "Error", err)
		exc = helperexception.Internal("failed commit transaction", err)
		return
	}

	// Output
	response.AdminID = entity.AdminID
	response.Name = entity.Name
	response.Email = entity.Email
	response.Username = entity.Username

	return
}

func (usecase *AdminUseCase) Select(ctx context.Context, request *internalmodeladminrequest.SelectAdminInfo) (response internalmodeladminresponse.SelectAdminResponse, exc *helperexception.Exception) {

	// Init Transaction
	tx := usecase.DB.WithContext(ctx)

	// Exec DB
	result, totalItems, page, pageSize, err := usecase.AdminRepository.Select(tx, ctx, request.QueryInfo)
	if err != nil {
		usecase.Log.Error("usecase.Select()", "AdminRepository.Select()", "Error", err)
		exc = helperexception.Internal("failed select admin", err)
		return
	}

	// Output
	response.Data = internalentity.ConvertEntitiesToModelResponse(result)
	response.TotalItems = totalItems
	response.TotalPages = page
	response.Page = page
	response.PageSize = pageSize

	return
}

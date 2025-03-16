package internaldeliveryhttpadmin

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	coreresponse "panel-ektensi/core/response"
	helpergenerator "panel-ektensi/helper/generator"
	internalmodeladminrequest "panel-ektensi/internal/model/admin/request"
	internalmodeladminresponse "panel-ektensi/internal/model/admin/response"
	internalusecaseadmin "panel-ektensi/internal/usecase/admin"
	"time"
)

type AdminHandler struct {
	UseCase *internalusecaseadmin.AdminUseCase
	Log     *slog.Logger
}

func NewAdminHandler(
	usecase *internalusecaseadmin.AdminUseCase,
	log *slog.Logger,
) *AdminHandler {
	return &AdminHandler{
		UseCase: usecase,
		Log:     log,
	}
}

func (handler *AdminHandler) Create(fiberCtx *fiber.Ctx) error {
	var timeIn = time.Now()

	// Init Ctx
	ctx := helpergenerator.DefaultContextGenerator(fiberCtx)

	// Get Request Body
	requestData := &internalmodeladminrequest.CreateAdminInfo{}
	if err := fiberCtx.BodyParser(&requestData); err != nil {
		handler.Log.Info("fiberCtx.BodyParser()", "Info", err)
		errString := err.Error()
		return fiberCtx.JSON(coreresponse.ApiResponse[internalmodeladminresponse.CreateAdminResponse]{
			Tin:     timeIn,
			Tout:    time.Now(),
			Success: false,
			Status:  fiber.StatusBadRequest,
			Error:   &errString,
			Latency: helpergenerator.GetLatency(timeIn),
			Data:    nil,
		})
	}

	// Exec UseCase
	response, exc := handler.UseCase.Create(ctx, requestData)
	if exc != nil {
		handler.Log.Info("handler.UseCase.Create()", "Info", exc)
		return fiberCtx.JSON(coreresponse.ApiResponse[internalmodeladminresponse.CreateAdminResponse]{
			Tin:     timeIn,
			Tout:    time.Now(),
			Success: false,
			Status:  exc.GetHttpCode(),
			Error:   exc.GetError(),
			Latency: helpergenerator.GetLatency(timeIn),
			Data:    nil,
		})
	}

	return fiberCtx.JSON(coreresponse.ApiResponse[internalmodeladminresponse.CreateAdminResponse]{
		Tin:     timeIn,
		Tout:    time.Now(),
		Success: false,
		Status:  fiber.StatusCreated,
		Error:   nil,
		Latency: helpergenerator.GetLatency(timeIn),
		Data:    &response,
	})
}

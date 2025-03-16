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

// Create
//
//	@Summary		Insert Admin.
//	@Description	Insert Admin.
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			data	body		internalmodeladminrequest.CreateAdminInfo		true	"Insert Order Request Parameter"
//	@Success		201		{object}	coreresponse.ApiResponse[internalmodeladminresponse.CreateAdminResponse]		"Result"
//	@Failure		400		{object}	coreresponse.ApiResponse[internalmodeladminresponse.CreateAdminResponse]		"Result"
//	@Router			/admin [post]
//	@Security		Bearer
func (handler *AdminHandler) Create(fiberCtx *fiber.Ctx) error {
	var timeIn = time.Now()

	// Init Ctx
	ctx := helpergenerator.DefaultContextGenerator(fiberCtx)

	// Get Request Body
	requestData := &internalmodeladminrequest.CreateAdminInfo{}
	if err := fiberCtx.BodyParser(&requestData); err != nil {
		errString := err.Error()
		handler.Log.Info("handler.Create()", "fiberCtx.BodyParser()", "Info", err)
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(coreresponse.ApiResponse[internalmodeladminresponse.CreateAdminResponse]{
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
		handler.Log.Info("handler.Create()", "UseCase.Create()", "Info", exc)
		return fiberCtx.Status(exc.GetHttpCode()).JSON(coreresponse.ApiResponse[internalmodeladminresponse.CreateAdminResponse]{
			Tin:     timeIn,
			Tout:    time.Now(),
			Success: false,
			Status:  exc.GetHttpCode(),
			Error:   exc.GetError(),
			Latency: helpergenerator.GetLatency(timeIn),
			Data:    nil,
		})
	}

	return fiberCtx.Status(fiber.StatusCreated).JSON(coreresponse.ApiResponse[internalmodeladminresponse.CreateAdminResponse]{
		Tin:     timeIn,
		Tout:    time.Now(),
		Success: true,
		Status:  fiber.StatusCreated,
		Error:   nil,
		Latency: helpergenerator.GetLatency(timeIn),
		Data:    &response,
	})
}

// Select
//
//	@Summary		Get All Admin.
//	@Description	Get All Admin.
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			filter		query		string											false						"Search Parameter"
//	@Param			sort		query		string											false						"Sorting Parameter"
//	@Param			page		query		int												false						"Current Page"
//	@Param			pageSize	query		int												false						"Rows Count"
//	@Success		200			{object}	coreresponse.ApiResponse[internalmodeladminresponse.SelectAdminResponse]		"Result"
//	@Router			/admin [get]
//	@Security		Bearer
func (handler *AdminHandler) Select(fiberCtx *fiber.Ctx) error {
	var timeIn = time.Now()

	// Init Ctx
	ctx := helpergenerator.DefaultContextGenerator(fiberCtx)

	// Get Param
	queryInfo, err := helpergenerator.GenerateQueryInfoPostgreSQL(fiberCtx)
	if err != nil {
		errString := err.Error()
		handler.Log.Info("handler.Select()", "helpergenerator.GenerateQueryInfoPostgreSQL()", "Info", err)
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(coreresponse.ApiResponse[internalmodeladminresponse.SelectAdminResponse]{
			Tin:     timeIn,
			Tout:    time.Now(),
			Success: false,
			Status:  fiber.StatusBadRequest,
			Error:   &errString,
			Latency: helpergenerator.GetLatency(timeIn),
			Data:    nil,
		})
	}

	// Pointing Param
	requestData := &internalmodeladminrequest.SelectAdminInfo{}
	requestData.QueryInfo = queryInfo

	// Exec UseCase
	response, exc := handler.UseCase.Select(ctx, requestData)
	if exc != nil {
		handler.Log.Info("handler.Select()", "UseCase.Select()", "Info", exc)
		return fiberCtx.Status(exc.GetHttpCode()).JSON(coreresponse.ApiResponse[internalmodeladminresponse.SelectAdminResponse]{
			Tin:     timeIn,
			Tout:    time.Now(),
			Success: false,
			Status:  exc.GetHttpCode(),
			Error:   exc.GetError(),
			Latency: helpergenerator.GetLatency(timeIn),
			Data:    nil,
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(coreresponse.ApiResponse[internalmodeladminresponse.SelectAdminResponse]{
		Tin:     timeIn,
		Tout:    time.Now(),
		Success: true,
		Status:  fiber.StatusOK,
		Error:   nil,
		Latency: helpergenerator.GetLatency(timeIn),
		Data:    &response,
	})
}

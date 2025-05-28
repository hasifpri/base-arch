package internalentity

import (
	helperconverter "panel-ektensi/helper/converter"
	helpergenerator "panel-ektensi/helper/generator"
	"panel-ektensi/internal/model/request"
	"panel-ektensi/internal/model/response"
)

type Admin struct {
	AdminID  int64  `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255);unique"`
	Email    string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255)"`
	CommonEntity
}

func (t *Admin) TableName() string {
	return "admin"
}

func ConvertModelToEntitiesAdmin(model *internalmodeladminrequest.CreateAdminInfo) *Admin {

	// Generate Username
	return &Admin{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
		Username: helpergenerator.UsernameGenerator(model.Email),
	}
}

func ConvertEntitiesToModelResponse(entities []Admin) []internalmodeladminresponse.AdminReadResponse {

	models := []internalmodeladminresponse.AdminReadResponse{}
	for _, entity := range entities {
		models = append(models, internalmodeladminresponse.AdminReadResponse{
			AdminID:   entity.AdminID,
			Email:     entity.Email,
			Username:  entity.Username,
			Name:      entity.Email,
			CreatedAt: helperconverter.ConvertTimeToString(&entity.CreatedAt),
		})
	}

	return models
}

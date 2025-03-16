package internalentity

import (
	helpergenerator "panel-ektensi/helper/generator"
	internalmodeladminrequest "panel-ektensi/internal/model/admin/request"
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

package internalmodeladminrequest

import "panel-ektensi/core"

type CreateAdminInfo struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type SelectAdminInfo struct {
	QueryInfo core.QueryInfo
}

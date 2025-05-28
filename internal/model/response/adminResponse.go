package internalmodeladminresponse

type CreateAdminResponse struct {
	AdminID  int64  `json:"admin_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AdminReadResponse struct {
	AdminID   int64  `json:"admin_id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
type SelectAdminResponse struct {
	Data       []AdminReadResponse `json:"list"`
	TotalItems int32               `json:"total_items"`
	TotalPages int32               `json:"total_pages"`
	Page       int32               `json:"page"`
	PageSize   int32               `json:"page_size"`
}

package internalmodeladminresponse

type CreateAdminResponse struct {
	AdminID  int64  `json:"admin_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

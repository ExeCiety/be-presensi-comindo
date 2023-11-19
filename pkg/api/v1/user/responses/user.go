package responses

import "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"

type UserForLoginResponse struct {
	Id       string         `json:"id"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Name     string         `json:"name"`
	Roles    []*models.Role `json:"roles"`
	Token    string         `json:"token"`
}

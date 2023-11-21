package requests

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofrs/uuid"
)

type GetUsers struct {
	Search  string   `query:"search"`
	RoleIds []string `query:"role_ids"`
	utils.PaginationRequest
}

type CreateUser struct {
	Id       uuid.UUID      `json:"-"`
	Username string         `json:"username" validate:"required,unique_login_username"`
	Email    string         `json:"email" validate:"required,email,unique_login_username"`
	Nik      string         `json:"nik" validate:"required,unique_login_username"`
	Password string         `json:"password" validate:"required"`
	Name     string         `json:"name" validate:"required"`
	RoleIds  []string       `json:"role_ids" validate:"required,min=1,dive,uuid" gorm:"-"`
	Roles    []*models.Role `json:"-"`
}

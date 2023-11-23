package requests

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"
)

type FindUsers struct {
	Search  string   `query:"search"`
	RoleIds []string `query:"role_ids"`
	utils.PaginationRequest
}

type FindUser struct {
	Identity string
}

type CreateUser struct {
	Username string   `json:"username" validate:"required,unique_login_username"`
	Email    string   `json:"email" validate:"required,email,unique_login_username"`
	Nik      string   `json:"nik" validate:"required,unique_login_username"`
	Password string   `json:"password" validate:"required"`
	Name     string   `json:"name" validate:"required"`
	RoleIds  []string `json:"role_ids" validate:"required,min=1,dive,uuid,exists=roles;id" gorm:"-"`
}

type UpdateUser struct {
	Identity string   `json:"-"`
	Username string   `json:"username" validate:"omitempty,unique_login_username"`
	Email    string   `json:"email" validate:"omitempty,email,unique_login_username"`
	Nik      string   `json:"nik" validate:"omitempty,unique_login_username"`
	Password string   `json:"password" validate:"omitempty"`
	Name     string   `json:"name" validate:"omitempty"`
	RoleIds  []string `json:"role_ids" validate:"omitempty,min=1,dive,uuid,exists=roles;id" gorm:"-"`
}

type DeleteUsers struct {
	Ids []string `json:"ids" validate:"required,min=1,dive,uuid,exists=users;id"`
}

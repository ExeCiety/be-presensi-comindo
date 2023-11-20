package requests

import "github.com/ExeCiety/be-presensi-comindo/utils"

type GetUsers struct {
	Search  string   `query:"search"`
	RoleIds []string `query:"role_ids"`
	utils.PaginationRequest
}

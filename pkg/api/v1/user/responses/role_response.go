package responses

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/enums"

	"github.com/gofrs/uuid"
)

type SimpleRole struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	RoleName string    `json:"role_name"`
}

type FullRole struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	RoleName    string    `json:"role_name"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

func (SimpleRole) TableName() string {
	return enums.RoleTableName
}

func (FullRole) TableName() string {
	return enums.RoleTableName
}

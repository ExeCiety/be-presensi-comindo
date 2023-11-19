package models

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofrs/uuid"
)

type RoleUser struct {
	Id     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	RoleId uuid.UUID `json:"role_id"`
	UserId uuid.UUID `json:"user_id"`
	utils.Timestamp
}

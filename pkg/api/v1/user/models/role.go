package models

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofrs/uuid"
)

type Role struct {
	Id          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name"`
	RoleName    string    `json:"role_name"`
	Description string    `json:"description"`
	utils.Timestamp
}

package models

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofrs/uuid"
)

type OvertimeActivity struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	PresenceId uuid.UUID `json:"presence_id"`
	Presence   *Presence `json:"presence" gorm:"foreignKey:PresenceId;references:Id"`
	Activity   string    `json:"activity"`
	utils.Timestamp
}

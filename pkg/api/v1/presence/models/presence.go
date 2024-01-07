package models

import (
	"time"

	userModels "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"

	"github.com/ExeCiety/be-presensi-comindo/utils"
	"github.com/gofrs/uuid"
)

type Presence struct {
	Id         uuid.UUID        `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserId     uuid.UUID        `json:"user_id"`
	User       *userModels.User `json:"user" gorm:"foreignKey:UserId;references:Id"`
	EntryTime  time.Time        `json:"entry_time"`
	ExitTime   *time.Time       `json:"exit_time"`
	IsOvertime *bool            `json:"is_overtime"`
	utils.Timestamp
}

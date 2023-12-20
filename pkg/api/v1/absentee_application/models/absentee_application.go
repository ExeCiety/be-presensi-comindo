package models

import (
	"time"

	userModels "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofrs/uuid"
)

type AbsenteeApplication struct {
	Id         uuid.UUID        `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserId     uuid.UUID        `json:"user_id"`
	User       *userModels.User `json:"user" gorm:"foreignKey:UserId;references:Id"`
	Type       string           `json:"type"`
	Status     string           `json:"status"`
	DateStart  time.Time        `json:"date_start"`
	DateEnd    time.Time        `json:"date_end"`
	Reason     *string          `json:"reason"`
	Attachment *string          `json:"attachment"`
	utils.Timestamp
}

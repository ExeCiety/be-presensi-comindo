package responses

import (
	"time"

	userResponses "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"

	"github.com/gofrs/uuid"
)

type GetAbsenteeApplications struct {
	Id         uuid.UUID                                      `json:"id"`
	UserId     uuid.UUID                                      `json:"-"`
	User       userResponses.UserForCreateAbsenteeApplication `json:"user" gorm:"foreignKey:UserId"`
	Type       string                                         `json:"type"`
	Status     string                                         `json:"status"`
	DateStart  time.Time                                      `json:"date_start"`
	DateEnd    time.Time                                      `json:"date_end"`
	Reason     string                                         `json:"reason"`
	Attachment string                                         `json:"attachment"`
	CreatedAt  string                                         `json:"created_at"`
	UpdatedAt  string                                         `json:"updated_at"`
}

type GetAbsenteeApplication struct {
	Id         uuid.UUID                                      `json:"id"`
	UserId     uuid.UUID                                      `json:"-"`
	User       userResponses.UserForCreateAbsenteeApplication `json:"user" gorm:"foreignKey:UserId"`
	Type       string                                         `json:"type"`
	Status     string                                         `json:"status"`
	DateStart  time.Time                                      `json:"date_start"`
	DateEnd    time.Time                                      `json:"date_end"`
	Reason     string                                         `json:"reason"`
	Attachment string                                         `json:"attachment"`
	CreatedAt  string                                         `json:"created_at"`
	UpdatedAt  string                                         `json:"updated_at"`
}

type CreateAbsenteeApplication struct {
	Id         uuid.UUID                                      `json:"id"`
	UserId     uuid.UUID                                      `json:"-"`
	User       userResponses.UserForCreateAbsenteeApplication `json:"user" gorm:"foreignKey:UserId"`
	Type       string                                         `json:"type"`
	Status     string                                         `json:"status"`
	DateStart  time.Time                                      `json:"date_start"`
	DateEnd    time.Time                                      `json:"date_end"`
	Reason     string                                         `json:"reason"`
	Attachment string                                         `json:"attachment"`
	CreatedAt  string                                         `json:"created_at"`
	UpdatedAt  string                                         `json:"updated_at"`
}

type UpdateAbsenteeApplication struct {
	Id         uuid.UUID
	UserId     uuid.UUID                                      `json:"-"`
	User       userResponses.UserForCreateAbsenteeApplication `json:"user" gorm:"foreignKey:UserId"`
	Type       string                                         `json:"type"`
	Status     string                                         `json:"status"`
	DateStart  time.Time                                      `json:"date_start"`
	DateEnd    time.Time                                      `json:"date_end"`
	Reason     string                                         `json:"reason"`
	Attachment string                                         `json:"attachment"`
	CreatedAt  string                                         `json:"created_at"`
	UpdatedAt  string                                         `json:"updated_at"`
}

type DeleteAbsenteeApplications struct {
	Id         uuid.UUID `json:"id"`
	Attachment string    `json:"attachment"`
}

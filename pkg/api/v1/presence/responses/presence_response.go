package responses

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/enums"
	userResponses "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"

	"github.com/gofrs/uuid"
)

type GetPresences struct {
	Id         string  `json:"id"`
	UserId     string  `json:"-"`
	EntryTime  string  `json:"entry_time"`
	ExitTime   *string `json:"exit_time"`
	IsOvertime bool    `json:"is_overtime"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`

	User               userResponses.UserForPresence          `json:"user" gorm:"foreignKey:UserId"`
	OvertimeActivities []CreateOvertimeActivitiesFromPresence `json:"overtime_activities" gorm:"foreignKey:PresenceId;references:Id"`
}

type GetPresence struct {
	Id         string  `json:"id"`
	UserId     string  `json:"-"`
	EntryTime  string  `json:"entry_time"`
	ExitTime   *string `json:"exit_time"`
	IsOvertime bool    `json:"is_overtime"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`

	User               userResponses.UserForPresence          `json:"user" gorm:"foreignKey:UserId"`
	OvertimeActivities []CreateOvertimeActivitiesFromPresence `json:"overtime_activities" gorm:"foreignKey:PresenceId;references:Id"`
}

type CreatePresence struct {
	Id         string  `json:"id"`
	UserId     string  `json:"-"`
	EntryTime  string  `json:"entry_time"`
	ExitTime   *string `json:"exit_time"`
	IsOvertime bool    `json:"is_overtime"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`

	User               *userResponses.UserForPresence         `json:"user" gorm:"foreignKey:UserId"`
	OvertimeActivities []CreateOvertimeActivitiesFromPresence `json:"overtime_activities" gorm:"foreignKey:PresenceId;references:Id"`
}

type UpdatePresence struct {
	Id         string  `json:"id"`
	UserId     string  `json:"-"`
	EntryTime  string  `json:"entry_time"`
	ExitTime   *string `json:"exit_time"`
	IsOvertime bool    `json:"is_overtime"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`

	User               userResponses.UserForPresence          `json:"user" gorm:"foreignKey:UserId"`
	OvertimeActivities []CreateOvertimeActivitiesFromPresence `json:"overtime_activities" gorm:"foreignKey:PresenceId;references:Id"`
}

type DeletePresences struct {
	Id uuid.UUID `json:"id"`
}

func (CreatePresence) TableName() string {
	return enums.PresenceTableName
}

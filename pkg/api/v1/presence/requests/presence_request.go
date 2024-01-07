package requests

import "github.com/ExeCiety/be-presensi-comindo/utils"

type GetPresences struct {
	UserId    *string `query:"user_id" validate:"omitempty,uuid,exists=users;id"`
	EntryTime *string `query:"entry_time" validate:"omitempty,datetime=2006-01-02"`
	ExitTime  *string `query:"exit_time" validate:"omitempty,datetime=2006-01-02"`
	utils.PaginationRequest
}

type GetPresence struct {
	Id     string  `query:"id"`
	UserId *string `query:"-"`
}

type CheckIfPresenceExistOnThatDay struct {
	UserId    string
	EntryTime string
	ExitTime  string
}

type CreatePresence struct {
	UserId       *string `json:"user_id" validate:"omitempty,uuid,exists=users;id"`
	EntryTime    string  `json:"entry_time" validate:"required,datetime=2006-01-02 15:04:05,date_same_as_today=2006-01-02 15:04:05"`
	ExitTime     *string `json:"exit_time" validate:"omitempty,datetime=2006-01-02 15:04:05,date_same_as_field=entry time;2006-01-02 15:04:05"`
	PresenceLat  float64 `json:"presence_lat" validate:"required,latitude"`
	PresenceLong float64 `json:"presence_long" validate:"required,longitude"`

	OvertimeActivities *[]CreateOvertimeActivitiesFromPresence `json:"overtime_activities"`
}

type UpdatePresence struct {
	Id           string  `json:"-"`
	EntryTime    *string `json:"entry_time"`
	ExitTime     string  `json:"exit_time"`
	PresenceLat  float64 `json:"presence_lat"`
	PresenceLong float64 `json:"presence_long"`

	OvertimeActivities *[]CreateOvertimeActivitiesFromPresence `json:"overtime_activities"`
}

type UpdatePresenceForAdmin struct {
	Id           string  `json:"-"`
	EntryTime    *string `json:"entry_time" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	ExitTime     string  `json:"exit_time" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	PresenceLat  float64 `json:"presence_lat" validate:"required,latitude"`
	PresenceLong float64 `json:"presence_long" validate:"required,longitude"`

	OvertimeActivities *[]CreateOvertimeActivitiesFromPresence `json:"overtime_activities"`
}

type UpdatePresenceForHrdAndEmployee struct {
	Id           string  `json:"-"`
	ExitTime     string  `json:"exit_time" validate:"required_with=EntryTime,omitempty,datetime=2006-01-02 15:04:05,date_same_as_today=2006-01-02 15:04:05"`
	PresenceLat  float64 `json:"presence_lat" validate:"required,latitude"`
	PresenceLong float64 `json:"presence_long" validate:"required,longitude"`

	OvertimeActivities *[]CreateOvertimeActivitiesFromPresence `json:"overtime_activities"`
}

type DeletePresences struct {
	Ids []string `json:"ids" validate:"required,min=1,dive,uuid,exists=presences;id"`
}

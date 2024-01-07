package responses

import "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/enums"

type CreateOvertimeActivitiesFromPresence struct {
	Id         string `json:"id"`
	PresenceId string `json:"presence_id"`
	Activity   string `json:"activity"`
}

type DeleteOverTimeActivitiesFromPresence struct {
	Id string `json:"id"`
}

func (CreateOvertimeActivitiesFromPresence) TableName() string {
	return enums.OvertimeTableName
}

func (DeleteOverTimeActivitiesFromPresence) TableName() string {
	return enums.OvertimeTableName
}

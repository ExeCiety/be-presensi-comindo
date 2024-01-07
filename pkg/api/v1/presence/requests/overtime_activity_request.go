package requests

type CreateOvertimeActivitiesFromPresence struct {
	Activity string `json:"activity" validate:"required"`
}

type DeleteOverTimeActivitiesFromPresence struct {
	PresenceId *string `json:"presence_id"`
}

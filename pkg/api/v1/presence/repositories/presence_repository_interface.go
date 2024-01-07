package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/responses"

	"gorm.io/gorm"
)

type PresenceRepositoryInterface interface {
	// Find Presences
	FindPresences(db *gorm.DB, request *requests.GetPresences, result *[]responses.GetPresences) error

	// Find Presence
	FindPresence(db *gorm.DB, request *requests.GetPresence, result *responses.GetPresence) error
	CheckIfPresenceExistOnThatDay(db *gorm.DB, request *requests.CheckIfPresenceExistOnThatDay) bool

	// Create Presence
	CreatePresence(db *gorm.DB, payload *models.Presence, result *responses.CreatePresence) error
	CreateOvertimeActivities(db *gorm.DB, payload *[]*models.OvertimeActivity, result *[]responses.CreateOvertimeActivitiesFromPresence) error

	// Update Presence
	UpdatePresence(db *gorm.DB, request *requests.UpdatePresence, payload *models.Presence, result *responses.UpdatePresence) error

	// Delete Presences
	DeletePresences(db *gorm.DB, request *requests.DeletePresences, result *[]responses.DeletePresences) error
	DeleteOvertimeActivities(db *gorm.DB, request *requests.DeleteOverTimeActivitiesFromPresence, result *[]responses.DeleteOverTimeActivitiesFromPresence) error
}

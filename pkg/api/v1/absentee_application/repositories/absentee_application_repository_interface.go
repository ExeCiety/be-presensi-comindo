package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/responses"

	"gorm.io/gorm"
)

type AbsenteeApplicationRepositoryInterface interface {
	// Find AbsenteeApplications
	FindAbsenteeApplications(db *gorm.DB, request *requests.GetAbsenteeApplications, result *[]responses.GetAbsenteeApplications) error

	// Find AbsenteeApplication
	FindAbsenteeApplication(db *gorm.DB, request *requests.GetAbsenteeApplication, result *responses.GetAbsenteeApplication) error
	CheckIfAbsenteeApplicationExistOnThatDays(db *gorm.DB, request *requests.CheckIfAbsenteeApplicationExistOnThatDays) bool

	// Create AbsenteeApplication
	CreateAbsenteeApplication(db *gorm.DB, payload *models.AbsenteeApplication, result *responses.CreateAbsenteeApplication) error

	// Update AbsenteeApplication
	UpdateAbsenteeApplication(db *gorm.DB, request *requests.UpdateAbsenteeApplication, payload *models.AbsenteeApplication, result *responses.UpdateAbsenteeApplication) error

	// Delete AbsenteeApplication
	DeleteAbsenteeApplications(db *gorm.DB, request *requests.DeleteAbsenteeApplications, result *[]responses.DeleteAbsenteeApplications) error
}

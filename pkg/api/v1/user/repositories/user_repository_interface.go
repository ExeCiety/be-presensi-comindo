package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	// Find Users
	FindUsers(db *gorm.DB, request *requests.FindUsers, result *[]responses.FindUsers) error

	// Find User
	FindUser(db *gorm.DB, request *requests.FindUser, result *responses.FindUser) error
	FindUserForLogin(db *gorm.DB, request *requests.FindUser, result *models.User) error
	IsUserByIdentityExist(db *gorm.DB, username string) bool

	// Create User
	CreateUser(db *gorm.DB, payload *models.User, result *responses.CreateUser) error

	// Update User
	UpdateUser(db *gorm.DB, request *requests.UpdateUser, payload *models.User, result *responses.UpdateUser) error

	// Delete Users
	DeleteUsers(db *gorm.DB, request *requests.DeleteUsers, response *[]responses.DeleteUsers) error
}

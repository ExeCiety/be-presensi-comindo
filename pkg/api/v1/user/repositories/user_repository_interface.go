package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	// Find Users
	FindUsers(db *gorm.DB, request *requests.GetUsers, result *[]responses.GetUsers) error

	// Find User
	FindUserByUsernameOrEmailOrNik(db *gorm.DB, username string, result *models.User) error
	IsUserByUsernameOrEmailOrNikExist(db *gorm.DB, username string) bool

	// Create User
	CreateUser(db *gorm.DB, payload *models.User, result *responses.CreateUser) error
}

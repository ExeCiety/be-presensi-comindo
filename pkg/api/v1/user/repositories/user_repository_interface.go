package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindUserByUsernameOrEmailOrNik(db *gorm.DB, username string, user *models.User) error
}

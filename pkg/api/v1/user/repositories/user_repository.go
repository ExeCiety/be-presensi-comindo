package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (lr *UserRepository) FindUserByUsernameOrEmailOrNik(db *gorm.DB, username string, user *models.User) error {
	tx := db.Model(models.User{}).
		Preload("Roles").
		Where("username = ?", username).
		Or("email = ?", username).
		Or("nik = ?", username)

	return tx.Find(&user).Error
}

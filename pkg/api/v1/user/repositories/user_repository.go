package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (lr *UserRepository) FindUsers(db *gorm.DB, request *requests.GetUsers, result *[]responses.GetUsers) error {
	tx := db.Model(models.User{})
	baseGetUsers(tx, request)

	return utils.Paginate(tx, &request.PaginationRequest).Find(&result).Error
}

func (lr *UserRepository) FindUserByUsernameOrEmailOrNik(db *gorm.DB, username string, result *models.User) error {
	tx := db.Model(models.User{}).
		Preload("Roles").
		Where("username = ?", username).
		Or("email = ?", username).
		Or("nik = ?", username)

	return tx.First(&result).Error
}

func baseGetUsers(tx *gorm.DB, request *requests.GetUsers) {
	tx.Preload("Roles").
		Order("created_at ASC")

	if request.Search != "" {
		tx.Where(
			"(name LIKE ? OR username LIKE ? OR email LIKE ? OR nik LIKE ?)",
			"%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%", "%"+request.Search+"%",
		)
	}

	if len(request.RoleIds) > 0 {
		tx.Where("(EXISTS (SELECT ru.user_id, ru.role_id FROM role_users AS ru WHERE ru.user_id = users.id AND ru.role_id IN (?)))", request.RoleIds)
	}
}

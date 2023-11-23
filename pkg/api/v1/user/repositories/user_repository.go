package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (ur *UserRepository) FindUsers(db *gorm.DB, request *requests.FindUsers, result *[]responses.FindUsers) error {
	tx := db.Model(models.User{}).Unscoped()
	baseGetUsers(tx, request)

	return utils.Paginate(tx, &request.PaginationRequest).
		Find(&result).Error
}

func (ur *UserRepository) FindUser(db *gorm.DB, request *requests.FindUser, result *responses.FindUser) error {
	tx := db.Model(models.User{}).
		Unscoped().
		Preload("Roles").
		Scopes(models.WhereByIdentity(request.Identity))

	return tx.First(&result).Error
}

func (ur *UserRepository) FindUserForLogin(
	db *gorm.DB,
	request *requests.FindUser,
	result *models.User,
) error {
	tx := db.Model(models.User{}).
		Preload("Roles").
		Scopes(models.WhereByIdentity(request.Identity))

	return tx.First(&result).Error
}

func (ur *UserRepository) IsUserByIdentityExist(db *gorm.DB, username string) bool {
	var countUser int64

	db.Model(models.User{}).
		Unscoped().
		Preload("Roles").
		Scopes(models.WhereByIdentity(username)).
		Count(&countUser)

	if countUser > 0 {
		return true
	}

	return false
}

func (ur *UserRepository) CreateUser(db *gorm.DB, payload *models.User, result *responses.CreateUser) error {
	if err := db.Model(models.User{}).Create(&payload).Error; err != nil {
		return err
	}

	return db.Model(models.User{}).
		Preload("Roles").
		First(&result, payload.Id).Error
}

func (ur *UserRepository) UpdateUser(
	db *gorm.DB,
	request *requests.UpdateUser,
	payload *models.User,
	result *responses.UpdateUser,
) error {
	if err := db.Model(models.User{}).Scopes(models.WhereByIdentity(request.Identity)).Updates(&payload).Error; err != nil {
		return err
	}

	return db.Model(models.User{}).
		Unscoped().
		Scopes(models.WhereByIdentity(request.Identity)).
		Preload("Roles").
		First(&result).Error
}

func (ur *UserRepository) DeleteUsers(db *gorm.DB, request *requests.DeleteUsers, response *[]responses.DeleteUsers) error {
	tx := db.Model(models.User{}).Unscoped()

	if len(request.Ids) > 0 {
		tx.Where("id IN (?)", request.Ids)
	}

	return tx.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&response).Error
}

func baseGetUsers(tx *gorm.DB, request *requests.FindUsers) {
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

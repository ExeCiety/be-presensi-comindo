package services

import (
	"github.com/ExeCiety/be-presensi-comindo/db"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/repositories"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db       *gorm.DB
	userRepo repositories.UserRepositoryInterface
}

func NewUserService(userRepositoryInterface repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		db:       db.DB,
		userRepo: userRepositoryInterface,
	}
}

func (us UserService) FindUsers(c *fiber.Ctx, request *requests.GetUsers, responseData *[]responses.GetUsers) error {
	if err := utils.QueryParserAndValidate(c, request); err != nil {
		return err
	}

	if err := us.userRepo.FindUsers(us.db, request, responseData); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	return nil
}

func (us UserService) CreateUser(c *fiber.Ctx, request *requests.CreateUser, response *responses.CreateUser) error {
	if err := utils.BodyParserAndValidate(c, request); err != nil {
		return err
	}

	// Set Password
	password, errHashPassword := utils.HashPassword(request.Password)
	if errHashPassword != nil {
		log.Error(errHashPassword)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	// Set Roles
	var roles []*models.Role
	for _, roleId := range request.RoleIds {
		roles = append(roles, &models.Role{Id: uuid.FromStringOrNil(roleId)})
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Nik:      request.Nik,
		Password: password,
		Name:     request.Name,
		Roles:    roles,
	}

	if err := us.userRepo.CreateUser(us.db, &user, response); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	return nil
}

package services

import (
	"github.com/ExeCiety/be-presensi-comindo/db"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/repositories"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

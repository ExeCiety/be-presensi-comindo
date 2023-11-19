package services

import (
	"time"

	"github.com/ExeCiety/be-presensi-comindo/db"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/requests"
	userModels "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	userRepositories "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/repositories"
	userResponses "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type LoginService struct {
	db       *gorm.DB
	userRepo userRepositories.UserRepositoryInterface
}

func NewLoginService(userRepositoryInterface userRepositories.UserRepositoryInterface) LoginServiceInterface {
	return &LoginService{
		db:       db.DB,
		userRepo: userRepositoryInterface,
	}
}

func (ls *LoginService) Login(
	c *fiber.Ctx,
	req *requests.LoginRequest,
	userForLogin *userResponses.UserForLoginResponse,
) error {
	if err := utils.BodyParserAndValidate(c, req); err != nil {
		return err
	}

	var user userModels.User

	if err := ls.userRepo.FindUserByUsernameOrEmailOrNik(ls.db, req.Username, &user); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError,
			utils.Translate("err.internal_server_error", nil),
			nil,
		)
	}

	if user.Id.IsNil() {
		errMsg := utils.Translate("err.incorrect_username_or_password", nil)
		log.Error(errMsg)
		return utils.NewApiError(fiber.StatusUnauthorized, errMsg, nil)
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusUnauthorized,
			utils.Translate("err.incorrect_username_or_password", nil),
			nil,
		)
	}

	var jwtToken string
	if err := generateJwtTokenForLogin(&user, &jwtToken); err != nil {
		return err
	}

	err := copier.Copy(userForLogin, &user)
	if err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}
	userForLogin.Token = jwtToken

	return nil
}

func generateJwtTokenForLogin(user *userModels.User, outputToken *string) error {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["nik"] = user.Nik
	claims["name"] = user.Name
	claims["roles"] = user.Roles

	expiryTime := utils.StrToInt(utils.GetEnvValue("JWT_EXPIRES_IN_HOURS", "1"))
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expiryTime)).Unix()

	token, err := jwtToken.SignedString([]byte(utils.GetEnvValue("JWT_SECRET", "secret")))
	if err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utils.Translate("err.internal_server_error", nil), nil,
		)
	}

	*outputToken = token

	return nil
}

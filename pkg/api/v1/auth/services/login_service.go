package services

import (
	"time"

	"github.com/ExeCiety/be-presensi-comindo/db"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/requests"
	userModels "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/models"
	userRepositories "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/repositories"
	userRequests "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	userResponses "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"
	utilsValidations "github.com/ExeCiety/be-presensi-comindo/utils/validations"

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
	request *requests.LoginRequest,
	responseData *userResponses.UserForLogin,
) error {
	if err := utilsValidations.BodyParserAndValidate(c, request); err != nil {
		return err
	}

	var user userModels.User
	if err := ls.userRepo.FindUserForLogin(ls.db, &userRequests.FindUser{Identity: request.Username}, &user); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	if user.Id.IsNil() {
		errMsg := utils.Translate("err.incorrect_username_or_password", nil)
		log.Error(errMsg)
		return utils.NewApiError(fiber.StatusUnauthorized, errMsg, nil)
	}

	if err := utils.ComparePassword(user.Password, request.Password); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusUnauthorized,
			utils.Translate("err.incorrect_username_or_password", nil),
			nil,
		)
	}

	var jwtToken string
	if err := generateJwtTokenForLogin(&user, &jwtToken, utilsEnums.JwtAccessTokenType); err != nil {
		return err
	}

	var jwtRefreshToken string
	if err := generateJwtTokenForLogin(&user, &jwtRefreshToken, utilsEnums.JwtRefreshTokenType); err != nil {
		return err
	}

	if err := copier.Copy(responseData, &user); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	responseData.Token = jwtToken
	responseData.RefreshToken = jwtRefreshToken

	return nil
}

func generateJwtTokenForLogin(user *userModels.User, outputToken *string, jwtType string) error {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["id"] = user.Id
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["nik"] = user.Nik
	claims["name"] = user.Name
	claims["roles"] = user.Roles

	jwtExpiresInHours := utils.StrToInt(utils.GetEnvValue("JWT_EXPIRES_IN_HOURS", "1"))
	if jwtType == utilsEnums.JwtRefreshTokenType {
		jwtExpiresInHours = utils.StrToInt(utils.GetEnvValue("JWT_REFRESH_EXPIRES_IN_HOURS", "24"))
	}

	claims["exp"] = time.Now().Add(time.Hour * time.Duration(jwtExpiresInHours)).Unix()

	token, err := jwtToken.SignedString([]byte(utils.GetEnvValue("JWT_SECRET", "secret")))
	if err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	*outputToken = token

	return nil
}

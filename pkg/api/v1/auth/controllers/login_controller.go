package controllers

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/services"
	userResponses "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	LoginService services.LoginServiceInterface
}

func NewLoginController(loginService services.LoginServiceInterface) *LoginController {
	return &LoginController{
		LoginService: loginService,
	}
}

func (lc *LoginController) Login(c *fiber.Ctx) error {
	loginRequest := new(requests.LoginRequest)
	responseData := new(userResponses.UserForLoginResponse)

	if err := lc.LoginService.Login(c, loginRequest, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK, utils.Translate("login_success", nil), responseData, nil,
	)
}

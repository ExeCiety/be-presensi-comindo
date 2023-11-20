package controllers

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/services"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) FindUsers(c *fiber.Ctx) error {
	request := new(requests.GetUsers)
	responseData := new([]responses.GetUsers)

	if err := uc.UserService.FindUsers(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("user.get_users_success", nil),
		utils.GetResourceResponseData(responseData, request.PaginationRequest),
		nil,
	)
}

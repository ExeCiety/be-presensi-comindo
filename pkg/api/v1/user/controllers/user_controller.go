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
	request := new(requests.FindUsers)
	responseData := new([]responses.FindUsers)

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

func (uc *UserController) FindUser(c *fiber.Ctx) error {
	request := new(requests.FindUser)
	responseData := new(responses.FindUser)

	if err := uc.UserService.FindUser(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("user.get_users_success", nil),
		responseData,
		nil,
	)
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	request := new(requests.CreateUser)
	responseData := new(responses.CreateUser)

	if err := uc.UserService.CreateUser(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusCreated,
		utils.Translate("user.create_user_success", nil),
		responseData,
		nil,
	)
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	request := new(requests.UpdateUser)
	responseData := new(responses.UpdateUser)

	if err := uc.UserService.UpdateUser(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("user.update_user_success", nil),
		responseData,
		nil,
	)
}

func (uc *UserController) DeleteUsers(c *fiber.Ctx) error {
	request := new(requests.DeleteUsers)
	responseData := new([]responses.DeleteUsers)

	if err := uc.UserService.DeleteUsers(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("user.delete_users_success", nil),
		responseData,
		nil,
	)
}

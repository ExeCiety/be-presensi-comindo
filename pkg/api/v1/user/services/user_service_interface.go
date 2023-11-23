package services

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"

	"github.com/gofiber/fiber/v2"
)

type UserServiceInterface interface {
	FindUsers(c *fiber.Ctx, request *requests.FindUsers, response *[]responses.FindUsers) error
	FindUser(c *fiber.Ctx, request *requests.FindUser, response *responses.FindUser) error
	CreateUser(c *fiber.Ctx, request *requests.CreateUser, response *responses.CreateUser) error
	UpdateUser(c *fiber.Ctx, request *requests.UpdateUser, response *responses.UpdateUser) error
	DeleteUsers(c *fiber.Ctx, request *requests.DeleteUsers, response *[]responses.DeleteUsers) error
}

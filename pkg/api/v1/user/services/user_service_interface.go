package services

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"

	"github.com/gofiber/fiber/v2"
)

type UserServiceInterface interface {
	FindUsers(c *fiber.Ctx, request *requests.GetUsers, response *[]responses.GetUsers) error
}

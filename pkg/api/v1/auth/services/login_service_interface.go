package services

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/requests"
	userResponses "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/responses"

	"github.com/gofiber/fiber/v2"
)

type LoginServiceInterface interface {
	Login(c *fiber.Ctx, request *requests.LoginRequest, responseData *userResponses.UserForLoginResponse) error
}

package middlewares

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
)

func ApiGuestMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString != "" {
		return c.Status(fiber.StatusForbidden).
			JSON(utils.ApiResponse{
				Message: utilsEnums.StatusMessageForbidden,
			})
	}

	return c.Next()
}

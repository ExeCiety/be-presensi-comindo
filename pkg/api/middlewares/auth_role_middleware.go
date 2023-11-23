package middlewares

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"
	"github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthRole protect routes
func AuthRole(roleName string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		if user == nil {
			return utils.SendApiResponse(c, fiber.StatusForbidden, enums.StatusMessageForbidden, nil, nil)
		}

		claims := user.Claims.(jwt.MapClaims)
		roles := claims["roles"].([]interface{})

		for _, role := range roles {
			if role.(map[string]interface{})["role_name"] == roleName {
				return c.Next()
			}
		}

		return utils.SendApiResponse(c, fiber.StatusForbidden, enums.StatusMessageForbidden, nil, nil)
	}
}

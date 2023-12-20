package middlewares

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsAuth "github.com/ExeCiety/be-presensi-comindo/utils/auth"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthRole protect routes
func AuthRole(roleName string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		if user == nil {
			return utils.SendApiResponse(c, fiber.StatusForbidden, utilsEnums.StatusMessageForbidden, nil, nil)
		}

		claims := user.Claims.(jwt.MapClaims)
		roles := claims["roles"].([]interface{})

		for _, role := range roles {
			if role.(map[string]interface{})["role_name"] == roleName {
				return c.Next()
			}
		}

		return utils.SendApiResponse(c, fiber.StatusForbidden, utilsEnums.StatusMessageForbidden, nil, nil)
	}
}

// AuthRoles protect routes
func AuthRoles(roleNames []string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		hasRoles, err := utilsAuth.IsUserAuthedHasRoles(roleNames)
		if err != nil {
			if err.Error() == utilsEnums.ErrorUserIsNotAuthenticated {
				return utils.SendApiResponse(c, fiber.StatusForbidden, utilsEnums.StatusMessageForbidden, nil, nil)
			}

			return utils.SendApiResponse(c, fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil, nil)
		}

		if hasRoles {
			return c.Next()
		}

		return utils.SendApiResponse(c, fiber.StatusForbidden, utilsEnums.StatusMessageForbidden, nil, nil)
	}
}

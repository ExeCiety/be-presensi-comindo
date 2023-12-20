package middlewares

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsAuth "github.com/ExeCiety/be-presensi-comindo/utils/auth"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// JwtAuth protect routes
func JwtAuth() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: utils.GetEnvValue("JWT_ALG", "HS256"),
			Key:    []byte(utils.GetEnvValue("JWT_SECRET", "secret")),
		},
		SuccessHandler: jwtSuccessHandler,
		ErrorHandler:   jwtError,
	})
}

func jwtSuccessHandler(c *fiber.Ctx) error {
	if err := utilsAuth.SetUserAuthed(c); err != nil {
		return err
	}

	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	log.Error(err)
	return c.Status(fiber.StatusUnauthorized).
		JSON(utils.ApiResponse{
			Message: utilsEnums.StatusMessageUnauthorized,
		})
}

package routers

import (
	apiV1AuthControllers "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	authRouter := router.Group("/auth")

	authRouter.Post("/login", apiV1AuthControllers.Login)
}

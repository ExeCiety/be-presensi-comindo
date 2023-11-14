package routers

import (
	apiV1AuthControllers "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	authRouter := router.Group("/auth")

	authRouter.Get("/login", apiV1AuthControllers.Login)
	authRouter.Get("/register", apiV1AuthControllers.Register)
}

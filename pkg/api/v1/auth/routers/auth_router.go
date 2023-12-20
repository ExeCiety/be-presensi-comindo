package routers

import (
	apiMiddlewares "github.com/ExeCiety/be-presensi-comindo/pkg/api/middlewares"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/controllers"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/services"
	userRepositories "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/repositories"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	userRepo := userRepositories.NewUserRepository()
	loginService := services.NewLoginService(userRepo)
	loginController := controllers.NewLoginController(loginService)

	authRouter := router.Group("/auth")
	authRouter.Post("/login", apiMiddlewares.ApiGuestMiddleware, loginController.Login)
}

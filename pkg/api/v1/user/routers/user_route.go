package routers

import (
	apiMiddlewares "github.com/ExeCiety/be-presensi-comindo/pkg/api/middlewares"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/controllers"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/enums"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/repositories"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/services"

	"github.com/gofiber/fiber/v2"
)

type AuthRouter struct {
}

func SetRouter(router fiber.Router) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userRouter := router.Group("/users")
	userRouter.Get("", apiMiddlewares.JwtAuth(), apiMiddlewares.AuthRole(enums.RoleNameAdmin), userController.FindUsers)
	userRouter.Get("/:id", apiMiddlewares.JwtAuth(), apiMiddlewares.AuthRole(enums.RoleNameAdmin), userController.FindUser)
	userRouter.Post("", apiMiddlewares.JwtAuth(), apiMiddlewares.AuthRole(enums.RoleNameAdmin), userController.CreateUser)
	userRouter.Patch("/:id", apiMiddlewares.JwtAuth(), apiMiddlewares.AuthRole(enums.RoleNameAdmin), userController.UpdateUser)
	userRouter.Delete("", apiMiddlewares.JwtAuth(), apiMiddlewares.AuthRole(enums.RoleNameAdmin), userController.DeleteUsers)
}

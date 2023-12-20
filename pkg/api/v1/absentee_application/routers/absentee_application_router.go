package routers

import (
	apiMiddlewares "github.com/ExeCiety/be-presensi-comindo/pkg/api/middlewares"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/controllers"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/repositories"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/services"
	userEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/enums"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	absenteeApplicationRepo := repositories.NewAbsenteeApplicationRepository()
	absenteeApplicationService := services.NewAbsenteeApplicationService(absenteeApplicationRepo)
	absenteeApplicationController := controllers.NewAbsenteeApplicationController(absenteeApplicationService)

	absenteeApplicationRouter := router.Group("/absentee-applications")
	absenteeApplicationRouter.Get(
		"",
		apiMiddlewares.JwtAuth(),
		apiMiddlewares.AuthRoles([]string{
			userEnums.RoleNameAdmin, userEnums.RoleNameHrd, userEnums.RoleNameEmployee,
		}),
		absenteeApplicationController.FindAbsenteeApplications,
	)
	absenteeApplicationRouter.Get(
		"/:id",
		apiMiddlewares.JwtAuth(),
		apiMiddlewares.AuthRoles([]string{
			userEnums.RoleNameAdmin, userEnums.RoleNameHrd, userEnums.RoleNameEmployee,
		}),
		absenteeApplicationController.FindAbsenteeApplication,
	)
	absenteeApplicationRouter.Post(
		"",
		apiMiddlewares.JwtAuth(),
		apiMiddlewares.AuthRoles([]string{
			userEnums.RoleNameAdmin, userEnums.RoleNameEmployee,
		}),
		absenteeApplicationController.CreateAbsenteeApplication,
	)
	absenteeApplicationRouter.Patch(
		"/:id",
		apiMiddlewares.JwtAuth(),
		apiMiddlewares.AuthRoles([]string{
			userEnums.RoleNameAdmin, userEnums.RoleNameHrd, userEnums.RoleNameEmployee,
		}),
		absenteeApplicationController.UpdateAbsenteeApplication,
	)
	absenteeApplicationRouter.Delete(
		"",
		apiMiddlewares.JwtAuth(),
		apiMiddlewares.AuthRoles([]string{
			userEnums.RoleNameAdmin, userEnums.RoleNameEmployee,
		}),
		absenteeApplicationController.DeleteAbsenteeApplications,
	)
}

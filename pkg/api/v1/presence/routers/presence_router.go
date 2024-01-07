package routers

import (
	apiMiddlewares "github.com/ExeCiety/be-presensi-comindo/pkg/api/middlewares"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/controllers"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/repositories"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/services"
	userEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/enums"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	presenceRepo := repositories.NewPresenceRepository()
	presenceService := services.NewPresenceService(presenceRepo)
	presenceController := controllers.NewPresenceController(presenceService)

	presenceRouter := router.Group("/presences")
	presenceRouter.Get("/", apiMiddlewares.JwtAuth(), presenceController.FindPresences)
	presenceRouter.Get("/:id", apiMiddlewares.JwtAuth(), presenceController.FindPresence)
	presenceRouter.Post("/", apiMiddlewares.JwtAuth(), presenceController.CreatePresence)
	presenceRouter.Put("/:id", apiMiddlewares.JwtAuth(), presenceController.UpdatePresence)
	presenceRouter.Delete(
		"/",
		apiMiddlewares.JwtAuth(),
		apiMiddlewares.AuthRoles([]string{
			userEnums.RoleNameAdmin,
		}),
		presenceController.DeletePresences,
	)
}

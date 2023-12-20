package routers

import (
	absenteeApplicationRouters "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/routers"
	apiV1AuthRouters "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/auth/routers"
	userRouters "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/routers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	v1Router := router.Group("/v1")

	apiV1AuthRouters.SetRouter(v1Router)
	userRouters.SetRouter(v1Router)
	absenteeApplicationRouters.SetRouter(v1Router)
}

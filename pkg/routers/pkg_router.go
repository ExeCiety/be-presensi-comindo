package routers

import (
	apiRouters "github.com/ExeCiety/be-presensi-comindo/pkg/api/routers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(app *fiber.App) {
	apiRouters.SetRouter(app)
}

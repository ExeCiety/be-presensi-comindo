package routers

import (
	apiMiddlewares "github.com/ExeCiety/be-presensi-comindo/pkg/api/middlewares"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/controllers"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/services"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router) {
	fileService := services.NewFileService()
	fileController := controllers.NewFileController(fileService)

	fileRouter := router.Group("/files")
	fileRouter.Post("/upload", apiMiddlewares.JwtAuth(), fileController.Upload)
}

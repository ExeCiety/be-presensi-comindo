package services

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/responses"

	"github.com/gofiber/fiber/v2"
)

type FileServiceInterface interface {
	Upload(c *fiber.Ctx, request *requests.UploadFile, responseData *[]responses.UploadFile) error
}

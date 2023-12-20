package controllers

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/responses"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/services"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofiber/fiber/v2"
)

type FileController struct {
	FileService services.FileServiceInterface
}

func NewFileController(
	fileService services.FileServiceInterface,
) *FileController {
	return &FileController{
		FileService: fileService,
	}
}

func (fc FileController) Upload(c *fiber.Ctx) error {
	request := new(requests.UploadFile)
	responseData := new([]responses.UploadFile)

	if err := fc.FileService.Upload(c, request, responseData); err != nil {
		return err
	}

	return utils.SendApiResponse(
		c, fiber.StatusOK,
		utils.Translate("file.upload_file_success", nil),
		responseData,
		nil,
	)
}

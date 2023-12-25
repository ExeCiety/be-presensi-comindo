package services

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/enums"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"
	utilsFile "github.com/ExeCiety/be-presensi-comindo/utils/file"
	utilsValidations "github.com/ExeCiety/be-presensi-comindo/utils/validations"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type FileService struct {
	//
}

func NewFileService() FileServiceInterface {
	return &FileService{}
}

func (f FileService) Upload(c *fiber.Ctx, request *requests.UploadFile, responseData *[]responses.UploadFile) error {
	if err := utilsValidations.BodyParserAndValidate(c, request); err != nil {
		return err
	}

	// Parse multipart form
	form, errParseMultipartForm := c.MultipartForm()
	if errParseMultipartForm != nil {
		return errParseMultipartForm
	}
	request.File = form.File["file"]

	uploadFileResults, errUploadFilesToStorage := utilsFile.UploadFilesToStorage(c, request, utilsEnums.StorageNameTemp)
	if errUploadFilesToStorage != nil {
		log.Error(errUploadFilesToStorage)

		switch errUploadFilesToStorage.(type) {
		case *utilsFile.UploadFileError:
			switch errUploadFilesToStorage.(*utilsFile.UploadFileError).Name {
			case enums.UploadFileErrorFileTooLarge,
				enums.UploadFileErrorFileMimeTypeNotAllowed:
				return utils.NewApiError(
					fiber.StatusUnprocessableEntity,
					utils.Translate("err.validation_error", nil),
					map[string]string{
						"file": errUploadFilesToStorage.Error(),
					},
				)
			}
		}

		return utils.NewApiError(
			fiber.StatusInternalServerError, errUploadFilesToStorage.Error(), nil,
		)
	}

	for _, uploadFileResult := range uploadFileResults {
		*responseData = append(*responseData, responses.UploadFile{
			Filename:    uploadFileResult.Filename,
			FileUrl:     uploadFileResult.FileUrl,
			StorageName: utilsEnums.StorageNameTemp,
		})
	}

	return nil
}

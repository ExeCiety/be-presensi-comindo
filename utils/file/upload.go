package file

import (
	"fmt"
	"mime/multipart"
	"os"
	"time"

	fileEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/enums"
	fileRequests "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/requests"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsStorage "github.com/ExeCiety/be-presensi-comindo/utils/storage"

	"github.com/gofiber/fiber/v2"
)

func UploadFilesToStorage(c *fiber.Ctx, request *fileRequests.UploadFile, storageName string) ([]*UploadFileResult, error) {
	storageType, ok := utilsStorage.StorageTypes[storageName]
	if !ok {
		return nil, NewUploadFileError(
			fileEnums.UploadFileErrorStorageTypeNotFound,
			-1,
			"storage type not found",
		)
	}

	if err := utilsStorage.CheckStorageDirectoryExist(storageType); err != nil {
		return nil, NewUploadFileError(
			fileEnums.UploadFileErrorCheckStorageDirectoryExistFailed,
			-1,
			"check storage directory exist failed",
		)
	}

	uploadFilePurpose, okUploadFilePurposes := UploadFilePurposes[request.Purpose]
	if !okUploadFilePurposes {
		return nil, NewUploadFileError(
			fileEnums.UploadFileErrorPurposeNotFound,
			-1,
			"upload file purpose not found",
		)
	}

	// Check file size
	for index, file := range request.File {
		if file.Size > uploadFilePurpose.MaxSize {
			return nil, NewUploadFileError(
				fileEnums.UploadFileErrorFileTooLarge,
				index,
				fmt.Sprintf("file[%d] size is too large, maximum file size is %d bytes", index, uploadFilePurpose.MaxSize),
			)
		}
	}

	// Check file mime type
	for index, file := range request.File {
		isFileMimeTypesAllowed := false
		for _, allowedMimeType := range uploadFilePurpose.AllowedMimeTypes {
			if file.Header["Content-Type"][0] == allowedMimeType {
				isFileMimeTypesAllowed = true
				break
			}
		}

		if !isFileMimeTypesAllowed {
			return nil, NewUploadFileError(
				fileEnums.UploadFileErrorFileMimeTypeNotAllowed,
				index,
				fmt.Sprintf("file[%d] mime type is not allowed, allowed mime types is %v", index, uploadFilePurpose.AllowedMimeTypes),
			)
		}
	}

	var errorUploadFile *UploadFileError
	var uploadFileResults []*UploadFileResult

	for index, file := range request.File {
		newFilename := GenerateRandomUploadFileName(file)
		uploadFileResults = append(uploadFileResults, &UploadFileResult{
			Filename: newFilename,
			FileUrl:  GetFileUrlFromStorage(newFilename, storageType),
		})

		if err := c.SaveFile(file, fmt.Sprintf("./"+utilsStorage.GetStorageDirectory(storageType)+"/%s", newFilename)); err != nil {
			errorUploadFile = NewUploadFileError(
				fileEnums.UploadFileErrorSaveFileFailed,
				index,
				"failed to save file to storage",
			)
			break
		}
	}

	if errorUploadFile != nil {
		for index, uploadFileResult := range uploadFileResults {
			fileInfo, err := os.Stat(fmt.Sprintf("./"+utilsStorage.GetStorageDirectory(storageType)+"/%s", uploadFileResult.Filename))
			if os.IsNotExist(err) {
				continue
			}

			if fileInfo != nil {
				if errRemove := os.Remove(fmt.Sprintf("./"+utilsStorage.GetStorageDirectory(storageType)+"/%s", uploadFileResult.Filename)); errRemove != nil {
					return nil, NewUploadFileError(
						fileEnums.UploadFileErrorDeleteFailedUploadedFile,
						index,
						"failed to delete uploaded file",
					)
				}
			}
		}

		return nil, errorUploadFile
	}

	return uploadFileResults, nil
}

func GenerateRandomUploadFileName(file *multipart.FileHeader) string {
	return utils.GenerateRandomString(12) + "_" + utils.IntToStr(int(time.Now().Unix())) + GetFileExtensionFromFilename(file.Filename)
}

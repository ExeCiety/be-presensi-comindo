package file

import (
	"io"
	"os"
	"path/filepath"

	fileEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/enums"
	utilsStorage "github.com/ExeCiety/be-presensi-comindo/utils/storage"

	"github.com/gofiber/fiber/v2/log"
)

func GetFileUrlFromFilename(filename string, storageNameDestination string) (string, error) {
	storage, ok := utilsStorage.Storages[storageNameDestination]
	if !ok {
		return "", NewUploadFileError(
			fileEnums.UploadFileErrorStorageNotFound,
			-1,
			"storage not found",
		)
	}

	return GetFileUrlFromStorage(filename, storage), nil
}

func AssignFilesToStorage(payload *[]AssignFileToStoragePayload) (*[]AssignFileToStorageResult, error) {
	var result []AssignFileToStorageResult
	var errAssignFileToModel *AssignFileToModelError

	for _, moveFileToStoragePayload := range *payload {
		// Check Source Storage Exist
		sourceStorage, okSourceStorage := utilsStorage.Storages[moveFileToStoragePayload.SourceStorageName]
		if !okSourceStorage {
			errAssignFileToModel = NewAssignFileToModelError(
				fileEnums.UploadFileErrorStorageNotFound,
				fileEnums.UploadFileErrorStorageNotFound,
			)
			break
		}

		// Check Destination Storage Exist
		destinationStorage, okDestinationStorage := utilsStorage.Storages[moveFileToStoragePayload.DestinationStorageName]
		if !okDestinationStorage {
			errAssignFileToModel = NewAssignFileToModelError(
				fileEnums.UploadFileErrorStorageNotFound,
				fileEnums.UploadFileErrorStorageNotFound,
			)
			break
		}

		if !okSourceStorage || !okDestinationStorage {
			continue
		}

		// Open Source File
		sourceFilePath := filepath.Join(
			utilsStorage.GetStorageDirectory(sourceStorage), moveFileToStoragePayload.Filename,
		)
		sourceFile, errOpenSourceFile := os.Open(sourceFilePath)
		if errOpenSourceFile != nil {
			log.Error(errOpenSourceFile)

			if err := sourceFile.Close(); err != nil {
				log.Error(err)
			}

			continue
		}

		// Create Destination File
		destinationFilePath := filepath.Join(
			utilsStorage.GetStorageDirectory(destinationStorage), moveFileToStoragePayload.Filename,
		)
		destinationFile, errCreateDestinationFile := os.Create(destinationFilePath)
		if errCreateDestinationFile != nil {
			log.Error(errCreateDestinationFile)
			if err := sourceFile.Close(); err != nil {
				log.Error(err)
			}

			if err := destinationFile.Close(); err != nil {
				log.Error(err)
			}

			errAssignFileToModel = NewAssignFileToModelError(
				fileEnums.AssignFileToModelErrorFileNotFound,
				fileEnums.AssignFileToModelErrorFileNotFound,
			)

			break
		}

		// Copy Source File to Destination File
		_, errCopyFile := io.Copy(destinationFile, sourceFile)
		if err := sourceFile.Close(); err != nil {
			log.Error(err)
			errAssignFileToModel = NewAssignFileToModelError(
				err.Error(),
				err.Error(),
			)

			break
		}
		if err := destinationFile.Close(); err != nil {
			log.Error(err)
			errAssignFileToModel = NewAssignFileToModelError(
				err.Error(),
				err.Error(),
			)

			break
		}

		if errCopyFile != nil {
			log.Error(errCopyFile)
			errAssignFileToModel = NewAssignFileToModelError(
				fileEnums.AssignFileToModelErrorFileNotFound,
				fileEnums.AssignFileToModelErrorFileNotFound,
			)
		}

		result = append(result, AssignFileToStorageResult{
			Filename:               moveFileToStoragePayload.Filename,
			SourcePath:             sourceFilePath,
			DestinationPath:        destinationFilePath,
			SourceStorageName:      moveFileToStoragePayload.SourceStorageName,
			DestinationStorageName: moveFileToStoragePayload.DestinationStorageName,
		})
	}

	// If error, remove all destination file
	if errAssignFileToModel != nil {
		for _, assignFileToStorageResult := range result {
			fileInfo, err := os.Stat(assignFileToStorageResult.DestinationPath)
			if os.IsNotExist(err) {
				continue
			}

			if fileInfo != nil {
				if errRemove := os.Remove(assignFileToStorageResult.DestinationPath); errRemove != nil {
					return nil, NewAssignFileToModelError(
						fileEnums.AssignFileToModelErrorFailedToDeleteFailedAssignedFile,
						fileEnums.AssignFileToModelErrorFailedToDeleteFailedAssignedFile,
					)
				}
			}
		}

		return nil, errAssignFileToModel
	}

	// If success, remove all source file
	for _, assignFileToStorageResult := range result {
		var errDeleteSourceFile error

		storage := utilsStorage.Storages[assignFileToStorageResult.SourceStorageName]
		if err := RemoveFileFromStorage(assignFileToStorageResult.Filename, storage); err != nil {
			errDeleteSourceFile = err
		}

		if errDeleteSourceFile != nil {
			return nil, errDeleteSourceFile
		}
	}

	return &result, nil
}

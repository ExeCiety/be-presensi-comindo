package file

import (
	"net/url"
	"path"
	"strings"

	fileEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/enums"
	utilsStorage "github.com/ExeCiety/be-presensi-comindo/utils/storage"

	"github.com/gofiber/fiber/v2/log"
)

func RemoveFileFromModel(payload *[]RemoveFileFromModelPayload) (*[]RemoveFileFromModelResult, error) {
	for _, file := range *payload {
		storage, errGetStorageFromFileUrl := GetStorageFromFileUrl(file.FileUrl)
		if errGetStorageFromFileUrl != nil {
			return nil, errGetStorageFromFileUrl
		}

		filename, errGetFilenameFromFileUrl := GetFilenameFromFileUrl(file.FileUrl)
		if errGetFilenameFromFileUrl != nil {
			return nil, errGetFilenameFromFileUrl
		}

		// Remove file from storage
		if errRemoveFileFromStorage := RemoveFileFromStorage(filename, *storage); errRemoveFileFromStorage != nil {
			return nil, NewRemoveFileToModelError(
				"",
				"",
			)
		}
	}

	return nil, nil
}

func GetStorageFromFileUrl(fileUrl string) (*utilsStorage.BaseStorage, error) {
	// Parse file url
	fileUrlParse, errParseFileUrl := url.Parse(fileUrl)
	if errParseFileUrl != nil {
		log.Error(errParseFileUrl)

		return nil, NewRemoveFileToModelError(
			fileEnums.RemoveFileToModelErrorParseFileUrlFailed,
			fileEnums.RemoveFileToModelErrorParseFileUrlFailed,
		)
	}

	// Get storage name from file url
	storageName := ""
	pathSegments := strings.Split(fileUrlParse.Path, "/")
	for _, pathSegment := range pathSegments {
		for _, storage := range utilsStorage.Storages {
			if pathSegment == storage.Name {
				storageName = storage.Name
			}
		}
	}

	if storageName == "" {
		log.Error(fileEnums.RemoveFileToModelErrorFailedToGetStorageNameFromFileUrl)

		return nil, NewRemoveFileToModelError(
			fileEnums.RemoveFileToModelErrorFailedToGetStorageNameFromFileUrl,
			fileEnums.RemoveFileToModelErrorFailedToGetStorageNameFromFileUrl,
		)
	}

	storage, ok := utilsStorage.Storages[storageName]
	if !ok {
		log.Error(fileEnums.RemoveFileToModelErrorFailedToGetStorageNameFromFileUrl)

		return &utilsStorage.BaseStorage{}, NewRemoveFileToModelError(
			fileEnums.RemoveFileToModelErrorFailedToGetStorageNameFromFileUrl,
			fileEnums.RemoveFileToModelErrorFailedToGetStorageNameFromFileUrl,
		)
	}

	return &storage, nil
}

func GetFilenameFromFileUrl(fileUrl string) (string, error) {
	fileUrlParse, errParseFileUrl := url.Parse(fileUrl)

	if errParseFileUrl != nil {
		log.Error(errParseFileUrl)

		return "", NewRemoveFileToModelError(
			fileEnums.RemoveFileToModelErrorParseFileUrlFailed,
			fileEnums.RemoveFileToModelErrorParseFileUrlFailed,
		)
	}

	return path.Base(fileUrlParse.Path), nil
}

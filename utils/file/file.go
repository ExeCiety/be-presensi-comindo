package file

import (
	"path/filepath"

	fileEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/enums"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsStorage "github.com/ExeCiety/be-presensi-comindo/utils/storage"
)

type BaseUploadFilePurpose struct {
	MaxSize          int64
	AllowedMimeTypes []string
}

type UploadFileError struct {
	Name    string
	Index   int
	Message string
}

type UploadFileResult struct {
	Filename string
	FileUrl  string
}

var (
	UploadFilePurposes = map[string]BaseUploadFilePurpose{
		fileEnums.UploadFilePurposeAbsenteeApplicationAttachment: {
			MaxSize:          1 * 1024 * 1024,
			AllowedMimeTypes: []string{"image/jpeg", "image/png"},
		},
	}
)

func (u UploadFileError) Error() string {
	return u.Message
}

func NewUploadFileError(name string, index int, message string) *UploadFileError {
	return &UploadFileError{
		Name:    name,
		Index:   index,
		Message: message,
	}
}

func GetFileExtensionFromFilename(fileName string) string {
	return filepath.Ext(fileName)
}

func GetFileUrlFromStorage(fileName string, storageType utilsStorage.BaseStorageType) string {
	return utils.GetFullAddress() + "/" + utilsStorage.GetStorageDirectory(storageType) + "/" + fileName
}

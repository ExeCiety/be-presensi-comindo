package file

import (
	"os"
	"path"
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

type AssignFileToStoragePayload struct {
	Filename               string
	SourceStorageName      string
	DestinationStorageName string
}

type AssignFileToStorageResult struct {
	Filename               string
	SourcePath             string
	DestinationPath        string
	SourceStorageName      string
	DestinationStorageName string
}

type AssignFileToModelError struct {
	Name    string
	Message string
}

type RemoveFileFromModelPayload struct {
	FileUrl string
}

type RemoveFileFromModelResult struct {
	FileUrl string
}

type RemoveFileToModelError struct {
	Name    string
	Message string
}

type RemoveFileError struct {
	Name    string
	Message string
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

func (u AssignFileToModelError) Error() string {
	return u.Message
}

func (u RemoveFileToModelError) Error() string {
	return u.Message
}

func (u RemoveFileError) Error() string {
	return u.Message
}

func NewUploadFileError(name string, index int, message string) *UploadFileError {
	return &UploadFileError{
		Name:    name,
		Index:   index,
		Message: message,
	}
}

func NewAssignFileToModelError(name string, message string) *AssignFileToModelError {
	return &AssignFileToModelError{
		Name:    name,
		Message: message,
	}
}

func NewRemoveFileToModelError(name string, message string) *RemoveFileToModelError {
	return &RemoveFileToModelError{
		Name:    name,
		Message: message,
	}
}

func NewRemoveFileError(name string, message string) *RemoveFileError {
	return &RemoveFileError{
		Name:    name,
		Message: message,
	}
}

func GetFileExtensionFromFilename(filename string) string {
	return filepath.Ext(filename)
}

func GetFileUrlFromStorage(filename string, storage utilsStorage.BaseStorage) string {
	return utils.GetFullAddress() + "/" + utilsStorage.GetStorageDirectory(storage) + "/" + filename
}

func RemoveFileFromStorage(filename string, storage utilsStorage.BaseStorage) error {
	filePath := path.Join(utilsStorage.GetStorageDirectory(storage), filename)

	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil
	}

	if fileInfo != nil {
		if errRemove := os.Remove(filePath); errRemove != nil {
			return NewRemoveFileError(
				fileEnums.FailedToRemoveFileFromStorage,
				fileEnums.FailedToRemoveFileFromStorage,
			)
		}
	}

	return nil
}

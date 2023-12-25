package enums

import "time"

const (
	UploadFilePurposeAbsenteeApplicationAttachment = "absentee_application_attachment"
)

const (
	FailedToRemoveFileFromStorage = "failed to remove file from storage"

	AssignFileToModelErrorFileNotFound                     = "file not found"
	AssignFileToModelErrorFailedToDeleteFailedAssignedFile = "failed to delete failed Assigned file"

	UploadFileErrorFileTooLarge                     = "file too large"
	UploadFileErrorFileMimeTypeNotAllowed           = "file mime type not allowed"
	UploadFileErrorStorageNotFound                  = "storage not found"
	UploadFileErrorPurposeNotFound                  = "purpose not found"
	UploadFileErrorCheckStorageDirectoryExistFailed = "check storage directory exist failed"
	UploadFileErrorSaveFileFailed                   = "save file failed"
	UploadFileErrorDeleteFailedUploadedFile         = "failed to delete uploaded failed file"

	RemoveFileToModelErrorParseFileUrlFailed                = "parse file url failed"
	RemoveFileToModelErrorFailedToGetStorageNameFromFileUrl = "failed to get storage name from file url"
)

const (
	ExpiredTempFileDuration = time.Hour * 24 * 1
)

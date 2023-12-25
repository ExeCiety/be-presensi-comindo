package requests

import "mime/multipart"

type UploadFile struct {
	File    []*multipart.FileHeader `form:"-"`
	Purpose string                  `form:"purpose" validate:"required,oneof=absentee_application_attachment"`
}

type AssignFileToModel struct {
	Filename    string `json:"filename" validate:"required"`
	StorageName string `json:"storage_name" validate:"required,oneof=local temp"`
}

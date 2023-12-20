package requests

import "mime/multipart"

type UploadFile struct {
	File    []*multipart.FileHeader `form:"-"`
	Purpose string                  `form:"purpose" validate:"required,oneof=absentee_application_attachment"`
}

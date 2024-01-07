package services

import (
	"time"

	"github.com/ExeCiety/be-presensi-comindo/db"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/enums"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/repositories"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/responses"
	userEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/enums"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsAuth "github.com/ExeCiety/be-presensi-comindo/utils/auth"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"
	utilsFile "github.com/ExeCiety/be-presensi-comindo/utils/file"
	utilsValidations "github.com/ExeCiety/be-presensi-comindo/utils/validations"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AbsenteeApplicationService struct {
	db                      *gorm.DB
	absenteeApplicationRepo repositories.AbsenteeApplicationRepositoryInterface
}

func NewAbsenteeApplicationService(
	absenteeApplicationRepositoryInterface repositories.AbsenteeApplicationRepositoryInterface,
) AbsenteeApplicationServiceInterface {
	return &AbsenteeApplicationService{
		db:                      db.DB,
		absenteeApplicationRepo: absenteeApplicationRepositoryInterface,
	}
}

func (a AbsenteeApplicationService) FindAbsenteeApplications(
	c *fiber.Ctx,
	request *requests.GetAbsenteeApplications,
	responseData *[]responses.GetAbsenteeApplications,
) error {
	if err := utilsValidations.QueryParserAndValidate(c, request); err != nil {
		return err
	}

	if utilsAuth.IsUserAuthed() {
		userHasEmployeeRole, err := utilsAuth.IsUserAuthedHasRoles([]string{userEnums.RoleNameEmployee})
		if err != nil {
			log.Error(err)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		if userHasEmployeeRole {
			request.UserId = utils.StrToPtr(utilsAuth.UserAuthedData.Id.String())
		}
	}

	if err := a.absenteeApplicationRepo.FindAbsenteeApplications(a.db, request, responseData); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	return nil
}

func (a AbsenteeApplicationService) FindAbsenteeApplication(
	c *fiber.Ctx,
	request *requests.GetAbsenteeApplication,
	responseData *responses.GetAbsenteeApplication,
) error {
	request.Id = c.Params("id", "")

	if utilsAuth.IsUserAuthed() {
		userHasEmployeeRole, err := utilsAuth.IsUserAuthedHasRoles([]string{userEnums.RoleNameEmployee})
		if err != nil {
			log.Error(err)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		if userHasEmployeeRole {
			request.UserId = utils.StrToPtr(utilsAuth.UserAuthedData.Id.String())
		}
	}

	if err := a.absenteeApplicationRepo.FindAbsenteeApplication(a.db, request, responseData); err != nil {
		if err.Error() == utilsEnums.GormErrorRecordNotFound {
			return utils.NewApiError(
				fiber.StatusNotFound, utils.Translate("err.record_not_found", nil), nil,
			)
		}

		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	return nil
}

func (a AbsenteeApplicationService) CreateAbsenteeApplication(
	c *fiber.Ctx,
	request *requests.CreateAbsenteeApplication,
	responseData *responses.CreateAbsenteeApplication,
) error {
	if err := utilsValidations.BodyParserAndValidate(c, request); err != nil {
		return err
	}
	userId := uuid.Must(uuid.FromString(request.UserId))

	if utilsAuth.IsUserAuthed() {
		userHasEmployeeRole, err := utilsAuth.IsUserAuthedHasRoles([]string{userEnums.RoleNameEmployee})
		if err != nil {
			log.Error(err)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		if userHasEmployeeRole {
			if userId != utilsAuth.UserAuthedData.Id {
				return utils.NewApiError(
					fiber.StatusForbidden, utilsEnums.StatusMessageForbidden, nil,
				)
			}

			request.Status = enums.AbsenteeApplicationStatusInReview
		}
	}

	// Check if absentee application already exist on that days
	if isExist := a.absenteeApplicationRepo.CheckIfAbsenteeApplicationExistOnThatDays(
		a.db,
		&requests.CheckIfAbsenteeApplicationExistOnThatDays{
			UserId:    userId,
			DateStart: request.DateStart,
			DateEnd:   request.DateEnd,
		},
	); isExist == true {
		return utils.NewApiError(
			fiber.StatusConflict,
			utils.Translate("absentee_application.absentee_application_already_exist_on_that_days", nil),
			nil,
		)
	}

	// Prepare date start and date end for create payload
	dateStart, errParseDateStart := time.Parse(utilsEnums.DefaultDateFormat, request.DateStart)
	if errParseDateStart != nil {
		log.Error(errParseDateStart)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	dateEnd, errParseDateEnd := time.Parse(utilsEnums.DefaultDateFormat, request.DateEnd)
	if errParseDateEnd != nil {
		log.Error(errParseDateEnd)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	// Prepare attachment for create payload
	var attachment string
	if request.Attachment != nil {
		attachmentFileUrl, errGetFileUrlFromFilename := utilsFile.GetFileUrlFromFilename(
			request.Attachment.Filename, utilsEnums.DefaultStorageName,
		)
		if errGetFileUrlFromFilename != nil {
			log.Error(errGetFileUrlFromFilename)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		attachment = attachmentFileUrl
	}

	payload := models.AbsenteeApplication{
		UserId:     userId,
		Type:       request.Type,
		DateStart:  dateStart,
		DateEnd:    dateEnd,
		Status:     request.Status,
		Reason:     utils.StrToPtr(request.Reason),
		Attachment: utils.StrToPtr(attachment),
	}

	switch request.Type {
	case enums.AbsenteeApplicationTypePermission, enums.AbsenteeApplicationTypePaidLeave:
		payload.Attachment = nil
		break
	case enums.AbsenteeApplicationTypeSick:
		payload.Reason = nil
	}

	if err := a.absenteeApplicationRepo.CreateAbsenteeApplication(a.db, &payload, responseData); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	// Assign attachment to storage
	go func() {
		if request.Attachment != nil {
			_, errAssignFileToStorage := utilsFile.AssignFilesToStorage(&[]utilsFile.AssignFileToStoragePayload{
				{
					Filename:               request.Attachment.Filename,
					SourceStorageName:      request.Attachment.StorageName,
					DestinationStorageName: utilsEnums.DefaultStorageName,
				},
			})
			if errAssignFileToStorage != nil {
				log.Error(errAssignFileToStorage)
			}
		}
	}()

	return nil
}

func (a AbsenteeApplicationService) UpdateAbsenteeApplication(
	c *fiber.Ctx,
	request *requests.UpdateAbsenteeApplication,
	responseData *responses.UpdateAbsenteeApplication,
) error {
	if err := utilsValidations.BodyParserAndValidate(c, request); err != nil {
		return err
	}
	request.Id = c.Params("id", "")
	requestGetAbsenteeApplication := requests.GetAbsenteeApplication{
		Id: request.Id,
	}

	if utilsAuth.IsUserAuthed() {
		userHasEmployeeRole, err := utilsAuth.IsUserAuthedHasRoles([]string{userEnums.RoleNameEmployee})
		if err != nil {
			log.Error(err)
			return utils.SendApiResponse(c, fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil, nil)
		}

		if userHasEmployeeRole {
			requestGetAbsenteeApplication.UserId = utils.StrToPtr(utilsAuth.UserAuthedData.Id.String())

			request.UserId = utilsAuth.UserAuthedData.Id.String()
			request.Status = ""
			request.Type = ""
		}
	}

	// Get Absentee Application
	var absenteeApplication responses.GetAbsenteeApplication

	if err := a.absenteeApplicationRepo.FindAbsenteeApplication(a.db, &requestGetAbsenteeApplication, &absenteeApplication); err != nil {
		if err.Error() == utilsEnums.GormErrorRecordNotFound {
			return utils.NewApiError(
				fiber.StatusNotFound, utils.Translate("err.record_not_found", nil), nil,
			)
		}

		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	// Prepare date start and date end for update payload
	dateStart, errParseDateStart := time.Parse(utilsEnums.DefaultDateFormat, request.DateStart)
	if errParseDateStart != nil {
		log.Error(errParseDateStart)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	dateEnd, errParseDateEnd := time.Parse(utilsEnums.DefaultDateFormat, request.DateEnd)
	if errParseDateEnd != nil {
		log.Error(errParseDateEnd)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	// Prepare attachment for create payload
	var attachment string
	if request.Attachment != nil {
		attachmentFileUrl, errGetFileUrlFromFilename := utilsFile.GetFileUrlFromFilename(
			request.Attachment.Filename, utilsEnums.DefaultStorageName,
		)
		if errGetFileUrlFromFilename != nil {
			log.Error(errGetFileUrlFromFilename)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		attachment = attachmentFileUrl
	}

	payload := models.AbsenteeApplication{
		UserId:     uuid.FromStringOrNil(request.UserId),
		Type:       request.Type,
		Status:     request.Status,
		DateStart:  dateStart,
		DateEnd:    dateEnd,
		Reason:     utils.StrToPtr(request.Reason),
		Attachment: utils.StrToPtr(attachment),
	}

	switch request.Type {
	case enums.AbsenteeApplicationTypePermission, enums.AbsenteeApplicationTypePaidLeave:
		payload.Attachment = nil
		break
	case enums.AbsenteeApplicationTypeSick:
		payload.Reason = nil
	}

	if err := a.absenteeApplicationRepo.UpdateAbsenteeApplication(a.db, request, &payload, responseData); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	// Update Attachment
	go func() {
		if request.Attachment != nil {
			// Assign attachment to storage
			_, errAssignFileToStorage := utilsFile.AssignFilesToStorage(&[]utilsFile.AssignFileToStoragePayload{
				{
					Filename:               request.Attachment.Filename,
					SourceStorageName:      request.Attachment.StorageName,
					DestinationStorageName: utilsEnums.DefaultStorageName,
				},
			})
			if errAssignFileToStorage != nil {
				log.Error(errAssignFileToStorage)
			}

			// Remove old attachment from storage
			removeFileFromModelPayload := []utilsFile.RemoveFileFromModelPayload{
				{
					FileUrl: absenteeApplication.Attachment,
				},
			}

			_, errRemoveFileFromModel := utilsFile.RemoveFilesFromModel(&removeFileFromModelPayload)
			if errRemoveFileFromModel != nil {
				log.Error(errRemoveFileFromModel)
			}
		}
	}()

	return nil
}

func (a AbsenteeApplicationService) DeleteAbsenteeApplications(
	c *fiber.Ctx,
	request *requests.DeleteAbsenteeApplications,
	responseData *[]responses.DeleteAbsenteeApplications,
) error {
	if err := utilsValidations.BodyParserAndValidate(c, request); err != nil {
		return err
	}

	if utilsAuth.IsUserAuthed() {
		userHasEmployeeRole, err := utilsAuth.IsUserAuthedHasRoles([]string{userEnums.RoleNameEmployee})
		if err != nil {
			log.Error(err)
			return utils.SendApiResponse(c, fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil, nil)
		}

		if userHasEmployeeRole {
			request.UserId = utilsAuth.UserAuthedData.Id.String()
			request.Status = enums.AbsenteeApplicationStatusInReview
		}
	}

	if err := a.absenteeApplicationRepo.DeleteAbsenteeApplications(a.db, request, responseData); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	go func() {
		// Remove attachment from storage
		var removeFileFromModelPayload []utilsFile.RemoveFileFromModelPayload
		for _, absenteeApplication := range *responseData {
			if absenteeApplication.Attachment != "" {
				removeFileFromModelPayload = append(removeFileFromModelPayload, utilsFile.RemoveFileFromModelPayload{
					FileUrl: absenteeApplication.Attachment,
				})
			}
		}

		_, errRemoveFileFromModel := utilsFile.RemoveFilesFromModel(&removeFileFromModelPayload)
		if errRemoveFileFromModel != nil {
			log.Error(errRemoveFileFromModel)
		}
	}()

	return nil
}

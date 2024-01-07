package services

import (
	"github.com/jinzhu/copier"
	"time"

	"github.com/ExeCiety/be-presensi-comindo/db"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/enums"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/repositories"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/responses"
	userEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/enums"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsAuth "github.com/ExeCiety/be-presensi-comindo/utils/auth"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"
	utilsValidations "github.com/ExeCiety/be-presensi-comindo/utils/validations"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type PresenceService struct {
	db           *gorm.DB
	presenceRepo repositories.PresenceRepositoryInterface
}

func NewPresenceService(repositoryInterface repositories.PresenceRepositoryInterface) PresenceServiceInterface {
	return &PresenceService{
		db:           db.DB,
		presenceRepo: repositoryInterface,
	}
}

func (fs PresenceService) FindPresences(
	c *fiber.Ctx,
	request *requests.GetPresences,
	responseData *[]responses.GetPresences,
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

	if err := fs.presenceRepo.FindPresences(fs.db, request, responseData); err != nil {
		return err
	}

	return nil
}

func (fs PresenceService) FindPresence(
	c *fiber.Ctx,
	request *requests.GetPresence,
	responseData *responses.GetPresence,
) error {
	request.Id = c.Params("id", "")

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

	if err := fs.presenceRepo.FindPresence(fs.db, request, responseData); err != nil {
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

func (fs PresenceService) CreatePresence(
	c *fiber.Ctx,
	request *requests.CreatePresence,
	responseData *responses.CreatePresence,
) error {
	if err := utilsValidations.BodyParserAndValidate(c, request); err != nil {
		return err
	}

	// Prepare data from authed user
	if utilsAuth.IsUserAuthed() {
		userHasEmployeeRole, errCheckEmployeeRole := utilsAuth.IsUserAuthedHasRoles([]string{userEnums.RoleNameEmployee})
		if errCheckEmployeeRole != nil {
			log.Error(errCheckEmployeeRole)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		userHasHrdRole, errCheckHrdRole := utilsAuth.IsUserAuthedHasRoles([]string{userEnums.RoleNameHrd})
		if errCheckHrdRole != nil {
			log.Error(errCheckHrdRole)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		if userHasEmployeeRole || userHasHrdRole {
			request.UserId = utils.StrToPtr(utilsAuth.UserAuthedData.Id.String())
			request.ExitTime = nil
			request.OvertimeActivities = nil

			// Check if presence is not outside distance from office
			if fs.isPresenceNotOutsideDistanceFromOffice(request.PresenceLat, request.PresenceLong) {
				return utils.NewApiError(
					fiber.StatusForbidden,
					utils.Translate("presence.outside_distance_office", nil),
					nil,
				)
			}
		} else if request.UserId == nil {
			request.UserId = utils.StrToPtr(utilsAuth.UserAuthedData.Id.String())
		}
	}

	// Prepare entry time and exit time for create payload
	entryTime, errParseEntryTime := time.Parse(utilsEnums.DefaultDateTimeFormat, request.EntryTime)
	if errParseEntryTime != nil {
		log.Error(errParseEntryTime)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	var exitTime *time.Time
	if request.ExitTime != nil {
		exitTimeParse, errParseExitTime := time.Parse(utilsEnums.DefaultDateTimeFormat, *request.ExitTime)
		if errParseExitTime != nil {
			log.Error(errParseExitTime)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		exitTime = &exitTimeParse
	}

	// Check if presence already exist on that day
	if isExist := fs.presenceRepo.CheckIfPresenceExistOnThatDay(
		fs.db,
		&requests.CheckIfPresenceExistOnThatDay{
			UserId:    *request.UserId,
			EntryTime: entryTime.Truncate(24 * time.Hour).Format(utilsEnums.DefaultDateTimeFormat),
			ExitTime: entryTime.AddDate(0, 0, 1).
				Truncate(24 * time.Hour).
				Format(utilsEnums.DefaultDateTimeFormat),
		},
	); isExist == true {
		return utils.NewApiError(
			fiber.StatusConflict,
			utils.Translate("presence.already_exist_on_this_day", nil),
			nil,
		)
	}

	// Prepare payload
	payload := models.Presence{
		UserId:     uuid.Must(uuid.FromString(*request.UserId)),
		EntryTime:  entryTime,
		ExitTime:   exitTime,
		IsOvertime: utils.BoolToPtr(fs.isPresenceFromRequestOverTime(entryTime, exitTime)),
	}

	// Create presence
	tx := fs.db.Begin()

	if err := fs.presenceRepo.CreatePresence(tx, &payload, responseData); err != nil {
		tx.Rollback()
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	// Create overtime activities
	if request.OvertimeActivities != nil {
		var overtimeActivities []*models.OvertimeActivity
		var overtimeActivitiesResponse []responses.CreateOvertimeActivitiesFromPresence

		for _, v := range *request.OvertimeActivities {
			overtimeActivities = append(overtimeActivities, &models.OvertimeActivity{
				PresenceId: uuid.Must(uuid.FromString(responseData.Id)),
				Activity:   v.Activity,
			})
		}

		if err := fs.presenceRepo.CreateOvertimeActivities(tx, &overtimeActivities, &overtimeActivitiesResponse); err != nil {
			tx.Rollback()
			log.Error(err)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		responseData.OvertimeActivities = overtimeActivitiesResponse
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	return nil
}

func (fs PresenceService) UpdatePresence(
	c *fiber.Ctx,
	request *requests.UpdatePresence,
	responseData *responses.UpdatePresence,
) error {
	var userId *string

	if utilsAuth.IsUserAuthed() {
		userHasAdminRole, errCheckEmployeeRole := utilsAuth.IsUserAuthedHasRoles([]string{userEnums.RoleNameAdmin})
		if errCheckEmployeeRole != nil {
			log.Error(errCheckEmployeeRole)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		if userHasAdminRole {
			roleRequest := new(requests.UpdatePresenceForAdmin)
			if err := utilsValidations.BodyParserAndValidate(c, roleRequest); err != nil {
				log.Error(err)
				return err
			}
			if err := copier.Copy(request, &roleRequest); err != nil {
				log.Error(err)
				return utils.NewApiError(
					fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
				)
			}
		} else {
			roleRequest := new(requests.UpdatePresenceForHrdAndEmployee)

			if err := utilsValidations.BodyParserAndValidate(c, roleRequest); err != nil {
				log.Error(err)
				return err
			}
			if err := copier.Copy(request, &roleRequest); err != nil {
				log.Error(err)
				return utils.NewApiError(
					fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
				)
			}

			request.EntryTime = nil
			userId = utils.StrToPtr(utilsAuth.UserAuthedData.Id.String())
		}
	}

	request.Id = c.Params("id", "")

	// Check if presence is not outside distance from office
	if fs.isPresenceNotOutsideDistanceFromOffice(request.PresenceLat, request.PresenceLong) {
		return utils.NewApiError(
			fiber.StatusForbidden,
			utils.Translate("presence.outside_distance_office", nil),
			nil,
		)
	}

	// Get Presence
	var presence responses.GetPresence
	if err := fs.presenceRepo.FindPresence(
		fs.db,
		&requests.GetPresence{
			Id:     request.Id,
			UserId: userId,
		},
		&presence,
	); err != nil {
		if err.Error() == utilsEnums.GormErrorRecordNotFound {
			return utils.NewApiError(
				fiber.StatusNotFound, utils.Translate("err.record_not_found", nil), nil,
			)
		}

		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	// Prepare entry time and exit time for create payload
	var entryTime *time.Time
	if request.EntryTime != nil {
		entryTimeParse, errParseEntryTime := time.Parse(utilsEnums.DefaultDateTimeFormat, *request.EntryTime)
		if errParseEntryTime != nil {
			log.Error(errParseEntryTime)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}
		entryTime = &entryTimeParse
	} else if presence.EntryTime != "" {
		entryTimeParse, errParseEntryTime := time.Parse(time.RFC3339, presence.EntryTime)
		if errParseEntryTime != nil {
			log.Error(errParseEntryTime)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		entryTime = &entryTimeParse
	}

	var exitTime *time.Time
	if request.ExitTime != "" {
		exitTimeParse, errParseExitTime := time.Parse(utilsEnums.DefaultDateTimeFormat, request.ExitTime)
		if errParseExitTime != nil {
			log.Error(errParseExitTime)
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		exitTime = &exitTimeParse
	}

	// Check if exit time is greater than entry time
	if entryTime != nil && exitTime != nil {
		if err := fs.isPresenceExitTimeIsGreaterThanEntryTime(entryTime, exitTime); err == false {
			log.Error(err)
			return utils.NewApiError(
				fiber.StatusForbidden,
				utils.Translate("presence.exit_time_must_greater_than_entry_time", nil),
				nil,
			)
		}
	}

	// Prepare Is Overtime
	var isOvertime bool
	if entryTime != nil && exitTime != nil {
		isOvertime = fs.isPresenceFromRequestOverTime(*entryTime, exitTime)
	}

	// Prepare payload
	payload := models.Presence{
		EntryTime:  *entryTime,
		ExitTime:   exitTime,
		IsOvertime: &isOvertime,
	}

	// Update presence
	tx := fs.db.Begin()

	if err := fs.presenceRepo.UpdatePresence(tx, request, &payload, responseData); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	// Update overtime activities
	if request.OvertimeActivities != nil && len(*request.OvertimeActivities) > 0 {
		// Delete overtime activities
		if err := fs.presenceRepo.DeleteOvertimeActivities(
			tx,
			&requests.DeleteOverTimeActivitiesFromPresence{
				PresenceId: &request.Id,
			},
			nil,
		); err != nil {
			log.Error(err)
			tx.Rollback()
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		// Recreate overtime activities
		var overtimeActivities []*models.OvertimeActivity
		var overtimeActivitiesResponse []responses.CreateOvertimeActivitiesFromPresence

		for _, v := range *request.OvertimeActivities {
			overtimeActivities = append(overtimeActivities, &models.OvertimeActivity{
				PresenceId: uuid.Must(uuid.FromString(responseData.Id)),
				Activity:   v.Activity,
			})
		}

		if err := fs.presenceRepo.CreateOvertimeActivities(tx, &overtimeActivities, &overtimeActivitiesResponse); err != nil {
			log.Error(err)
			tx.Rollback()
			return utils.NewApiError(
				fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
			)
		}

		responseData.OvertimeActivities = overtimeActivitiesResponse
	}

	if err := tx.Commit().Error; err != nil {
		log.Error(err)
		tx.Rollback()
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	return nil
}

func (fs PresenceService) DeletePresences(
	c *fiber.Ctx,
	request *requests.DeletePresences,
	responseData *[]responses.DeletePresences,
) error {
	if err := utilsValidations.BodyParserAndValidate(c, request); err != nil {
		return err
	}

	if err := fs.presenceRepo.DeletePresences(fs.db, request, responseData); err != nil {
		log.Error(err)
		return utils.NewApiError(
			fiber.StatusInternalServerError, utilsEnums.StatusMessageInternalServerError, nil,
		)
	}

	return nil
}

func (fs PresenceService) isPresenceFromRequestOverTime(entryTime time.Time, exitTime *time.Time) bool {
	if exitTime == nil {
		return false
	}

	return exitTime.Sub(entryTime).Hours() >= enums.WorkingHoursPerDay+1
}

func (fs PresenceService) isPresenceNotOutsideDistanceFromOffice(presenceLat float64, presenceLong float64) bool {
	if utils.HaversineDistance(enums.OfficeLatitude, enums.OfficeLongitude, presenceLat, presenceLong) > enums.PresenceDistanceRadiusTolerance {
		return true
	}

	return false
}

func (fs PresenceService) isPresenceExitTimeIsGreaterThanEntryTime(entryTime *time.Time, exitTime *time.Time) bool {
	if entryTime == nil || exitTime == nil {
		return false
	}

	return exitTime.Sub(*entryTime).Seconds() > 0
}

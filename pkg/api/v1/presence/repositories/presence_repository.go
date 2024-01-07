package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/presence/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	"github.com/gofrs/uuid"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type PresenceRepository struct{}

func NewPresenceRepository() PresenceRepositoryInterface {
	return &PresenceRepository{}
}

func (pr PresenceRepository) FindPresences(
	db *gorm.DB,
	request *requests.GetPresences,
	result *[]responses.GetPresences,
) error {
	tx := db.Model(&models.Presence{}).Unscoped()
	baseFindPresences(tx, request)

	return utils.Paginate(tx, &request.PaginationRequest).
		Find(&result).Error
}

func (pr PresenceRepository) FindPresence(
	db *gorm.DB,
	request *requests.GetPresence,
	result *responses.GetPresence,
) error {
	tx := db.Model(&models.Presence{}).Unscoped()
	baseFindPresence(tx, request)

	return tx.First(&result).Error
}

func (pr PresenceRepository) CheckIfPresenceExistOnThatDay(
	db *gorm.DB,
	request *requests.CheckIfPresenceExistOnThatDay,
) bool {
	var dataCount int64

	tx := db.Model(models.Presence{}).
		Where("user_id = ?", request.UserId).
		Where("(entry_time >= ? AND entry_time <= ?)", request.EntryTime, request.ExitTime)

	tx.Count(&dataCount)

	if dataCount > 0 {
		return true
	}

	return false
}

func (pr PresenceRepository) CreatePresence(
	db *gorm.DB,
	payload *models.Presence,
	result *responses.CreatePresence,
) error {
	if err := db.Model(models.Presence{}).Create(&payload).Error; err != nil {
		return err
	}

	return db.Model(models.Presence{}).
		Preload("User.Roles").
		Preload("OvertimeActivities").
		First(&result, payload.Id).Error
}

func (pr PresenceRepository) CreateOvertimeActivities(
	db *gorm.DB,
	payload *[]*models.OvertimeActivity,
	result *[]responses.CreateOvertimeActivitiesFromPresence,
) error {
	if err := db.Model(models.OvertimeActivity{}).Create(&payload).Error; err != nil {
		return err
	}

	var ids []uuid.UUID
	for _, v := range *payload {
		ids = append(ids, v.Id)
	}

	return db.Model(models.OvertimeActivity{}).
		Where("id IN ?", ids).
		Find(&result).Error
}

func (pr PresenceRepository) UpdatePresence(
	db *gorm.DB,
	request *requests.UpdatePresence,
	payload *models.Presence,
	result *responses.UpdatePresence,
) error {
	tx := db.Model(models.Presence{}).Unscoped()
	baseFindPresence(tx, &requests.GetPresence{Id: request.Id})

	if err := tx.Updates(&payload).Error; err != nil {
		return err
	}

	return db.Model(models.Presence{}).
		Unscoped().
		Preload("User.Roles").
		Preload("OvertimeActivities").
		Where("id::varchar ILIKE ?", request.Id).
		First(&result).Error
}

func (pr PresenceRepository) DeletePresences(
	db *gorm.DB,
	request *requests.DeletePresences,
	result *[]responses.DeletePresences,
) error {
	tx := db.Model(&models.Presence{}).Unscoped()

	if len(request.Ids) > 0 {
		tx.Where("id IN (?)", request.Ids)
	}

	return tx.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&result).Error
}

func (pr PresenceRepository) DeleteOvertimeActivities(
	db *gorm.DB,
	request *requests.DeleteOverTimeActivitiesFromPresence,
	result *[]responses.DeleteOverTimeActivitiesFromPresence,
) error {
	tx := db.Model(&models.OvertimeActivity{}).Unscoped()

	if request.PresenceId != nil {
		tx.Where("presence_id = ?", request.PresenceId)
	}

	return tx.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&result).Error
}

func baseFindPresences(tx *gorm.DB, request *requests.GetPresences) {
	tx.Preload("User.Roles").
		Preload("OvertimeActivities")

	if request.UserId != nil {
		tx.Where("user_id = ?", request.UserId)
	}

	if request.EntryTime != nil && request.ExitTime == nil {
		tx.Where("entry_time >= ? OR exit_time >= ?", request.EntryTime, request.EntryTime)
	}

	if request.ExitTime != nil && request.EntryTime == nil {
		tx.Where("exit_time <= ? OR entry_time <= ?", request.ExitTime, request.ExitTime)
	}

	if request.EntryTime != nil && request.ExitTime != nil {
		tx.Where(
			"((entry_time >= ? AND entry_time <= ?) OR (exit_time >= ? AND exit_time <= ?))",
			request.EntryTime, request.ExitTime, request.EntryTime, request.ExitTime,
		)
	}
}

func baseFindPresence(tx *gorm.DB, request *requests.GetPresence) {
	tx.Preload("User.Roles").
		Preload("OvertimeActivities").
		Where("id::varchar ILIKE ?", request.Id)

	if request.UserId != nil {
		tx.Where("user_id = ?", request.UserId)
	}
}

package repositories

import (
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/models"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/requests"
	"github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/absentee_application/responses"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AbsenteeApplicationRepository struct{}

func NewAbsenteeApplicationRepository() AbsenteeApplicationRepositoryInterface {
	return &AbsenteeApplicationRepository{}
}

func (a AbsenteeApplicationRepository) FindAbsenteeApplications(
	db *gorm.DB,
	request *requests.GetAbsenteeApplications,
	result *[]responses.GetAbsenteeApplications,
) error {
	tx := db.Model(models.AbsenteeApplication{}).Unscoped()
	baseFindAbsenteeApplications(tx, request)

	return utils.Paginate(tx, &request.PaginationRequest).
		Find(&result).Error
}

func (a AbsenteeApplicationRepository) FindAbsenteeApplication(
	db *gorm.DB,
	request *requests.GetAbsenteeApplication,
	result *responses.GetAbsenteeApplication,
) error {
	tx := db.Model(models.AbsenteeApplication{}).Unscoped()

	baseFindAbsenteeApplication(tx, request)

	return tx.First(&result).Error
}

func (a AbsenteeApplicationRepository) CheckIfAbsenteeApplicationExistOnThatDays(
	db *gorm.DB,
	request *requests.CheckIfAbsenteeApplicationExistOnThatDays,
) bool {
	var dataCount int64

	tx := db.Model(models.AbsenteeApplication{}).
		Where("user_id = ?", request.UserId).
		Where("((date_start <= ? AND date_end >= ?) OR (date_start <= ? AND date_end >= ?))", request.DateStart, request.DateStart, request.DateEnd, request.DateEnd)

	if request.ExceptionId != nil {
		tx = tx.Where("id != ?", request.ExceptionId)
	}

	tx.Count(&dataCount)

	if dataCount > 0 {
		return true
	}

	return false
}

func (a AbsenteeApplicationRepository) CreateAbsenteeApplication(
	db *gorm.DB,
	payload *models.AbsenteeApplication,
	result *responses.CreateAbsenteeApplication,
) error {
	if err := db.Model(models.AbsenteeApplication{}).Create(&payload).Error; err != nil {
		return err
	}

	return db.Model(models.AbsenteeApplication{}).
		Preload("User.Roles").
		First(&result, payload.Id).Error
}

func (a AbsenteeApplicationRepository) UpdateAbsenteeApplication(
	db *gorm.DB,
	request *requests.UpdateAbsenteeApplication,
	payload *models.AbsenteeApplication,
	result *responses.UpdateAbsenteeApplication,
) error {
	if err := db.Model(models.AbsenteeApplication{}).Unscoped().Where("id::varchar ILIKE ?", request.Id).Updates(&payload).Error; err != nil {
		return err
	}

	return db.Model(models.AbsenteeApplication{}).
		Unscoped().
		Where("id::varchar ILIKE ?", request.Id).
		Preload("User.Roles").
		First(&result).Error
}

func (a AbsenteeApplicationRepository) DeleteAbsenteeApplications(
	db *gorm.DB,
	request *requests.DeleteAbsenteeApplications,
	result *[]responses.DeleteAbsenteeApplications,
) error {
	tx := db.Model(models.AbsenteeApplication{}).Unscoped()

	if len(request.Ids) > 0 {
		tx.Where("id IN (?)", request.Ids)
	}

	if request.UserId != "" {
		tx.Where("user_id = ?", request.UserId)
	}

	if request.Status != "" {
		tx.Where("status = ?", request.Status)
	}

	return tx.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}, {Name: "attachment"}}}).
		Delete(&result).Error
}

func baseFindAbsenteeApplications(tx *gorm.DB, request *requests.GetAbsenteeApplications) {
	tx.Preload("User.Roles")

	if request.UserId != nil {
		tx.Where("user_id = ?", request.UserId)
	}

	if request.DateStart != nil && request.DateEnd == nil {
		tx.Where("date_start >= ? OR date_end >= ?", request.DateStart, request.DateStart)
	}

	if request.DateEnd != nil && request.DateStart == nil {
		tx.Where("date_end <= ? OR date_start <= ?", request.DateEnd, request.DateEnd)
	}

	if request.DateStart != nil && request.DateEnd != nil {
		tx.Where(
			"((date_start >= ? AND date_start <= ?) OR (date_end >= ? AND date_end <= ?))",
			request.DateStart, request.DateEnd, request.DateStart, request.DateEnd,
		)
	}
}

func baseFindAbsenteeApplication(tx *gorm.DB, request *requests.GetAbsenteeApplication) {
	tx.Preload("User.Roles").
		Where("id::varchar ILIKE ?", request.Id)

	if request.UserId != nil {
		tx.Where("user_id = ?", request.UserId)
	}
}

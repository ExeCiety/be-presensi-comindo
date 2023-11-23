package validations

import (
	"strings"

	"github.com/ExeCiety/be-presensi-comindo/db"
	userRepositories "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterCustomValidations(v *validator.Validate) {
	if err := v.RegisterValidation("unique_login_username", UniqueLoginUsernameValidation); err != nil {
		log.Debug(err)
		return
	}

	if err := v.RegisterValidation("exists", Exists); err != nil {
		log.Debug(err)
		return
	}
}

func UniqueLoginUsernameValidation(fl validator.FieldLevel) bool {
	return !userRepositories.NewUserRepository().IsUserByIdentityExist(db.DB, fl.Field().String())
}

func Exists(fl validator.FieldLevel) bool {
	paramValues := strings.SplitN(fl.Param(), ";", 2)
	if len(paramValues) <= 0 {
		return false
	}

	tableName := paramValues[0]
	var columnName string

	if len(paramValues) > 1 {
		columnName = paramValues[1]
	} else {
		columnName = "id"
	}

	var resultCount int64
	db.DB.Table(tableName).Where(columnName+" = ?", fl.Field().String()).Count(&resultCount)

	if resultCount > 0 {
		return true
	}

	return false
}

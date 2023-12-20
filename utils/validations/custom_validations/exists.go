package custom_validations

import (
	"strings"

	"github.com/ExeCiety/be-presensi-comindo/db"

	"github.com/go-playground/validator/v10"
)

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

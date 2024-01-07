package custom_validations

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

func DateSameAsToday(fl validator.FieldLevel) bool {
	layout := fl.Param()
	if layout == "" {
		return false
	}

	fieldParse, errParseField := time.Parse(layout, fl.Field().String())
	if errParseField != nil {
		log.Error(errParseField)
		return false
	}

	today := utils.CreateTimeToday()
	if fieldParse.Format("2006-01-02") != today.Format("2006-01-02") {
		return false
	}

	return true
}

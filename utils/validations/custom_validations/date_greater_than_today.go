package custom_validations

import (
	"time"

	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
)

func DateGreaterThanToday(fl validator.FieldLevel) bool {
	layout := fl.Param()
	if layout == "" {
		return false
	}

	parse, err := time.Parse(layout, fl.Field().String())
	if err != nil {
		log.Error(err)
		return false
	}

	today := utils.CreateTimeToday()
	if parse.Before(today) {
		return false
	}

	return true
}

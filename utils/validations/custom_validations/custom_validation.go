package custom_validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterCustomValidations(v *validator.Validate) {
	if err := v.RegisterValidation("unique_login_username", UniqueLoginUsername); err != nil {
		log.Error(err)
		return
	}

	if err := v.RegisterValidation("exists", Exists); err != nil {
		log.Error(err)
		return
	}

	if err := v.RegisterValidation("date_greater_than_today", DateGreaterThanToday); err != nil {
		log.Error(err)
		return
	}

	if err := v.RegisterValidation("date_greater_than_field", DateGreaterThanField); err != nil {
		log.Error(err)
		return
	}

	if err := v.RegisterValidation("date_same_as_field", DateSameAsField); err != nil {
		log.Error(err)
		return
	}

	if err := v.RegisterValidation("date_same_as_today", DateSameAsToday); err != nil {
		log.Error(err)
		return
	}
}

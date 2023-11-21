package validations

import (
	"github.com/ExeCiety/be-presensi-comindo/db"
	userRepositories "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/repositories"
	"github.com/gofiber/fiber/v2/log"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations(v *validator.Validate) {
	if err := v.RegisterValidation("unique_login_username", UniqueLoginUsernameValidation); err != nil {
		log.Debug(err)
		return
	}
}

func UniqueLoginUsernameValidation(fl validator.FieldLevel) bool {
	return !userRepositories.NewUserRepository().IsUserByUsernameOrEmailOrNikExist(db.DB, fl.Field().String())
}

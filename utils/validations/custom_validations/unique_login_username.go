package custom_validations

import (
	"github.com/ExeCiety/be-presensi-comindo/db"
	userRepositories "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/user/repositories"

	"github.com/go-playground/validator/v10"
)

func UniqueLoginUsername(fl validator.FieldLevel) bool {
	return !userRepositories.NewUserRepository().
		IsUserByIdentityExist(db.DB, fl.Field().String())
}

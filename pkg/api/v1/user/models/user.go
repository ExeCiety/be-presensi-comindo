package models

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Nik         string    `json:"nik"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	utils.Timestamp

	Roles []*Role `json:"roles" gorm:"many2many:role_users;foreignKey:Id;joinForeignKey:UserId;References:Id;joinReferences:RoleId"`
}

func WhereByIdentity(identity string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id::varchar ILIKE ?", identity).
			Or("username = ?", identity).
			Or("email = ?", identity).
			Or("nik = ?", identity)
	}
}

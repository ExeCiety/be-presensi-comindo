package migrations

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Users struct {
	Id          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username    string    `json:"username" gorm:"type:varchar(20);not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null"`
	Nik         string    `json:"nik" gorm:"type:varchar(20);not null"`
	Password    string    `json:"password" gorm:"type:varchar(255); not null"`
	Name        string    `json:"name" gorm:"type:varchar(255); not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(20);"`
	utils.Timestamp
}

func CreateUsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231119053125",
		Migrate: func(tx *gorm.DB) error {
			return tx.Debug().
				AutoMigrate(&Users{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Debug().
				Migrator().DropTable(&Users{})
		},
	}
}

package migrations

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Roles struct {
	Id          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"type:varchar(20);not null"`
	RoleName    string    `json:"role_name" gorm:"type:varchar(20);not null"`
	Description string    `json:"description" gorm:"type:varchar(255);not null"`
	utils.Timestamp
}

func CreateRolesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231119055615",
		Migrate: func(tx *gorm.DB) error {
			return tx.Debug().
				AutoMigrate(&Roles{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Debug().
				Migrator().DropTable(&Roles{})
		},
	}
}

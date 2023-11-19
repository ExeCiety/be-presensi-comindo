package migrations

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"
	"github.com/go-gormigrate/gormigrate/v2"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type RoleUser struct {
	Id     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	RoleId uuid.UUID `json:"role_id"`
	Role   Roles     `gorm:"references:RoleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId uuid.UUID `json:"user_id"`
	User   Users     `gorm:"references:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	utils.Timestamp
}

func CreateRoleUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231119055621",
		Migrate: func(tx *gorm.DB) error {
			return tx.Debug().
				AutoMigrate(&RoleUser{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Debug().
				Migrator().DropTable(&RoleUser{})
		},
	}
}

package migrations

import (
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type OvertimeActivities struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	PresenceId uuid.UUID `json:"presence_id" gorm:"not null"`
	Presence   Presences `gorm:"references:PresenceId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Activity   string    `json:"activity" gorm:"type:text;not null"`
	utils.Timestamp
}

func CreateOvertimeActivities() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231231103958",
		Migrate: func(tx *gorm.DB) error {
			return tx.Debug().
				AutoMigrate(&OvertimeActivities{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Debug().
				Migrator().
				DropTable(&OvertimeActivities{})
		},
	}
}

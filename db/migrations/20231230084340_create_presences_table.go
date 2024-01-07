package migrations

import (
	"time"

	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Presences struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserId     uuid.UUID `json:"user_id" gorm:"not null"`
	User       Users     `gorm:"references:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EntryTime  time.Time `json:"entry_time" gorm:"not null"`
	ExitTime   time.Time `json:"exit_time" gorm:"default:null"`
	IsOvertime bool      `json:"is_overtime"`
	utils.Timestamp
}

func CreatePresencesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231230084340",
		Migrate: func(tx *gorm.DB) error {
			return tx.Debug().
				AutoMigrate(&Presences{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Debug().
				Migrator().
				DropTable(&Presences{})
		},
	}
}

package migrations

import (
	"time"

	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AbsenteeApplications struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserId     uuid.UUID `json:"user_id"`
	User       Users     `gorm:"references:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type       string    `json:"type" gorm:"type:varchar(20);not null"`
	Status     string    `json:"status" gorm:"type:varchar(20);default:'in_review'"`
	DateStart  time.Time `json:"date_start" gorm:"type:date;not null"`
	DateEnd    time.Time `json:"date_end" gorm:"type:date;not null"`
	Reason     string    `json:"reason" gorm:"type:text;"`
	Attachment string    `json:"attachment" gorm:"type:text;"`
	utils.Timestamp
}

func CreateAbsenteeApplicationsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231126143717",
		Migrate: func(tx *gorm.DB) error {
			return tx.Debug().
				AutoMigrate(&AbsenteeApplications{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Debug().
				Migrator().
				DropTable(&AbsenteeApplications{})
		},
	}
}

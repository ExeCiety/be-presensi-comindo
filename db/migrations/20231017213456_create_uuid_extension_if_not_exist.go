package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateUUIDExtensionIfNotExist() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231017213456",
		Migrate: func(tx *gorm.DB) error {
			return tx.Debug().
				Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Debug().
				Exec("DROP EXTENSION IF EXISTS \"uuid-ossp\";").Error
		},
	}
}

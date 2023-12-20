package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrations []*gormigrate.Migration

func init() {
	migrations = []*gormigrate.Migration{
		CreateUUIDExtensionIfNotExist(),
		CreateUsersTable(),
		CreateRolesTable(),
		CreateRoleUserTable(),
		CreateAbsenteeApplicationsTable(),
	}
}

func Execute(db *gorm.DB) error {
	m := gormigrate.New(db, &gormigrate.Options{
		TableName:                 "migrations",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: false,
	}, migrations)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}

	log.Printf("Migration did run successfully")
	return nil
}

func Rollback(db *gorm.DB) error {
	m := gormigrate.New(db, &gormigrate.Options{
		TableName:                 "migrations",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: false,
	}, migrations)

	if err := m.RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
		return err
	}

	log.Printf("Rolled back succesfully")
	return nil
}

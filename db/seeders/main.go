package seeders

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var seeders []*gormigrate.Migration

func init() {
	seeders = []*gormigrate.Migration{
		RoleSeeder(),
		UserSeeder(),
	}
}

func Execute(db *gorm.DB) error {
	m := gormigrate.New(db, &gormigrate.Options{
		TableName:                 "seeders",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: false,
	}, seeders)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}

	log.Printf("Seedering did run successfully")
	return nil
}

func Rollback(db *gorm.DB) error {
	m := gormigrate.New(db, &gormigrate.Options{
		TableName:                 "seeders",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: false,
	}, seeders)

	if err := m.RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
		return err
	}

	log.Printf("Seeder Rolled back succesfully")
	return nil
}

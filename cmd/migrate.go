package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ExeCiety/be-presensi-comindo/db"
	"github.com/ExeCiety/be-presensi-comindo/db/migrations"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/spf13/cobra"
)

var migrationTemplate = `package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type table struct {}

func migrationsTitle() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "migrationsID",
		Migrate: func(tx *gorm.DB) error {			
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}
`

// migrateCmd represents the migrating schema
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This command will migrate all migrations func inside the migrations folder.",
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("create")
		if fileName != "" {
			if err := createMigrationFile(fileName); err != nil {
				log.Fatalf("Error creating migration file: %v", err)
				return
			}
			fmt.Printf("Migration file %s created!\n", fileName)
			os.Exit(0)
		}

		db.Connect()

		if len(args) > 0 && args[0] == "rollback" {
			step, _ := cmd.Flags().GetString("step")
			if step == "" {
				step = "1"
			}

			for i := 0; i < utils.StrToInt(step); i++ {
				if err := migrations.Rollback(db.DB); err != nil {
					panic(err)
				}
			}
			fmt.Println("---Rollback success---")
			os.Exit(0)
		}

		if err := migrations.Execute(db.DB); err != nil {
			panic(err)
		}

		fmt.Println("---Migration success---")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.PersistentFlags().String("create", "", "Create migration file")
	migrateCmd.PersistentFlags().String("step", "", "Step rollback")
}

func createMigrationFile(filename string) error {
	rootPath := utils.GetRootPath()
	t := time.Now()
	currTime := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fileName := fmt.Sprintf("%s/db/migrations/%s_%s.go", rootPath, currTime, filename)

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating migration file: %v", err)
		return err
	}

	if _, err := f.WriteString(strings.Replace(migrationTemplate, "migrationsID", currTime, -1)); err != nil {
		log.Fatalf("Error creating migration file: %v", err)
		return err
	}

	if err := f.Close(); err != nil {
		log.Fatalf("Error creating migration file: %v", err)
		return err
	}

	return nil
}

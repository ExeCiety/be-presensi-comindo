package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ExeCiety/be-presensi-comindo/db"
	"github.com/ExeCiety/be-presensi-comindo/db/seeders"
	"github.com/ExeCiety/be-presensi-comindo/utils"

	"github.com/spf13/cobra"
)

var seederTemplate = `package seeders

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func seederTitle() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "seederID",
		Migrate: func(tx *gorm.DB) error {			
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}
`

// seederCmd represents the seeder schema
var seederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "This command will seed all seeder func inside the seeders folder.",
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("create")
		if fileName != "" {
			if err := createSeederFile(fileName); err != nil {
				log.Fatalf("Error creating seeder file: %v", err)
				return
			}
			fmt.Printf("Seeder file %s created!\n", fileName)
			os.Exit(0)
		}

		db.Connect()

		if len(args) > 0 && args[0] == "rollback" {
			step, _ := cmd.Flags().GetString("step")
			if step == "" {
				step = "1"
			}

			for i := 0; i < utils.StrToInt(step); i++ {
				if err := seeders.Rollback(db.DB); err != nil {
					panic(err)
				}
			}
			fmt.Println("---Rollback success---")
			os.Exit(0)
		}

		if err := seeders.Execute(db.DB); err != nil {
			panic(err)
		}

		fmt.Println("---Seedling success---")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(seederCmd)
	seederCmd.PersistentFlags().String("create", "", "Create seeder file")
	seederCmd.PersistentFlags().String("step", "", "Step rollback")
}

func createSeederFile(filename string) error {
	rootPath := utils.GetRootPath()
	t := time.Now()
	currTime := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fileName := fmt.Sprintf("%s/db/seeders/%s_%s.go", rootPath, currTime, filename)

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating seeder file: %v", err)
		return err
	}

	if _, err := f.WriteString(strings.Replace(seederTemplate, "seederID", currTime, -1)); err != nil {
		log.Fatalf("Error creating seeder file: %v", err)
		return err
	}

	if err := f.Close(); err != nil {
		log.Fatalf("Error creating seeder file: %v", err)
		return err
	}

	return nil
}

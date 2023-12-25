package jobs

import (
	"os"
	"path/filepath"
	"time"

	fileEnums "github.com/ExeCiety/be-presensi-comindo/pkg/api/v1/file/enums"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"
	utilsStorage "github.com/ExeCiety/be-presensi-comindo/utils/storage"

	"github.com/gofiber/fiber/v2/log"
)

// JobRemoveExpiredTempFiles is a cron job to remove expired temp files every 00:00:00
// Expired temp files are files that are not used for more than 1 day
func JobRemoveExpiredTempFiles() {
	_, err := MyCron.AddFunc("0 0 0 * * *", func() {
		log.Info("Running JobRemoveExpiredTempFiles")

		tempStorage := utilsStorage.Storages[utilsEnums.StorageNameTemp]
		folderPath := utilsStorage.GetStorageDirectory(tempStorage)

		err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				if time.Now().Sub(info.ModTime()) > fileEnums.ExpiredTempFileDuration {
					if errRemoveFile := os.Remove(path); errRemoveFile != nil {
						log.Errorf("Error deleting file %s: %v\n", path, errRemoveFile)
					}
				}
			}

			return nil
		})

		if err != nil {
			log.Error("Error walking the directory:", err)
		}
	})

	if err != nil {
		log.Error(err)

		return
	}
}

package storage

import (
	"os"

	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
)

type BaseStorageType struct {
	Name          string
	Type          string
	DirectoryName string
}

var (
	StorageTypes = map[string]BaseStorageType{
		utilsEnums.StorageNameLocal: {
			Name:          utilsEnums.StorageNameLocal,
			Type:          utilsEnums.StorageTypeLocal,
			DirectoryName: utilsEnums.StorageNameLocal,
		},
	}
)

func GetStorageDirectory(storageType BaseStorageType) string {
	baseDirectory := ""
	if storageType.Type == utilsEnums.StorageTypeLocal {
		baseDirectory = utilsEnums.BaseStorageTypeLocalPath + "/"
	}

	return baseDirectory + storageType.DirectoryName
}

func CheckStorageDirectoryExist(storageType BaseStorageType) error {
	if _, err := os.Stat("storage"); os.IsNotExist(err) {
		if errMkDirLog := os.Mkdir("storage", 0755); errMkDirLog != nil {
			return errMkDirLog
		}
	}

	if err := MakeStorageDirectoryByName(storageType); err != nil {
		return err
	}

	return nil
}

func MakeStorageDirectoryByName(storageType BaseStorageType) error {
	if _, err := os.Stat(GetStorageDirectory(storageType)); os.IsNotExist(err) {
		if errMkDirLog := os.Mkdir(GetStorageDirectory(storageType), 0755); errMkDirLog != nil {
			return errMkDirLog
		}
	}

	return nil
}

func RegisterStorageType(app *fiber.App) {
	for _, storageType := range StorageTypes {
		if storageType.Type == utilsEnums.StorageTypeLocal {
			storageDirectory := GetStorageDirectory(storageType)
			app.Static("/"+storageDirectory, "./"+storageDirectory)
		}
	}
}

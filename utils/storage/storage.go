package storage

import (
	"os"

	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

	"github.com/gofiber/fiber/v2"
)

type BaseStorage struct {
	Name          string
	Type          string
	DirectoryName string
}

var (
	Storages = map[string]BaseStorage{
		utilsEnums.StorageNameLocal: {
			Name:          utilsEnums.StorageNameLocal,
			Type:          utilsEnums.StorageTypeLocal,
			DirectoryName: utilsEnums.StorageNameLocal,
		},
		utilsEnums.StorageNameTemp: {
			Name:          utilsEnums.StorageNameTemp,
			Type:          utilsEnums.StorageTypeLocal,
			DirectoryName: utilsEnums.StorageNameTemp,
		},
	}
)

func GetStorageDirectory(storage BaseStorage) string {
	baseDirectory := ""
	if storage.Type == utilsEnums.StorageTypeLocal {
		baseDirectory = utilsEnums.BaseStorageTypeLocalPath + "/"
	}

	return baseDirectory + storage.DirectoryName
}

func CheckStorageDirectoryExist(storage BaseStorage) error {
	if _, err := os.Stat("storage"); os.IsNotExist(err) {
		if errMkDirLog := os.Mkdir("storage", 0755); errMkDirLog != nil {
			return errMkDirLog
		}
	}

	if err := MakeStorageDirectoryByName(storage); err != nil {
		return err
	}

	return nil
}

func MakeStorageDirectoryByName(storage BaseStorage) error {
	if _, err := os.Stat(GetStorageDirectory(storage)); os.IsNotExist(err) {
		if errMkDirLog := os.Mkdir(GetStorageDirectory(storage), 0755); errMkDirLog != nil {
			return errMkDirLog
		}
	}

	return nil
}

func RegisterStorages(app *fiber.App) {
	for _, storage := range Storages {
		if storage.Type == utilsEnums.StorageTypeLocal {
			storageDirectory := GetStorageDirectory(storage)
			app.Static("/"+storageDirectory, "./"+storageDirectory)
		}
	}
}

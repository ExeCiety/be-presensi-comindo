package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ExeCiety/be-presensi-comindo/cmd"
	"github.com/ExeCiety/be-presensi-comindo/db"
	pkgRouters "github.com/ExeCiety/be-presensi-comindo/pkg/routers"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"
	utilsStorage "github.com/ExeCiety/be-presensi-comindo/utils/storage"
	utilsValidations "github.com/ExeCiety/be-presensi-comindo/utils/validations"
	customValidations "github.com/ExeCiety/be-presensi-comindo/utils/validations/custom_validations"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Set root path
	cwd, _ := os.Getwd()
	utils.SetRootPath(cwd + "/")

	// Init Viper
	utils.InitViper()

	// Execute cmd tool
	cmd.Execute()

	// Connect DB
	db.Connect()

	// Set i18n
	utils.SetAllI18nBundles()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.DefaultErrorHandler,
	})

	// Set Helmet
	app.Use(helmet.New())

	// Set Recover
	app.Use(fiberRecover.New())

	// Configure Logger
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if errMkDirLog := os.Mkdir("logs", 0755); errMkDirLog != nil {
			panic(errMkDirLog)
		}
	}

	logFilePath := filepath.Join("logs", "logs.log")
	logFile, setLogError := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if setLogError != nil {
		panic(setLogError)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Failed to close log file")
		}
	}(logFile)

	iw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(iw)

	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Output: logFile,
	}))

	// Set Storage Directories
	if _, err := os.Stat("storage"); os.IsNotExist(err) {
		if errMkDirLog := os.Mkdir("storage", 0755); errMkDirLog != nil {
			panic(errMkDirLog)
		}
	}

	for _, storageType := range utilsStorage.StorageTypes {
		if err := utilsStorage.MakeStorageDirectoryByName(storageType); err != nil {
			panic(err)
		}
	}

	utilsStorage.RegisterStorageType(app)

	// Set Validator
	utilsValidations.MyValidation = validator.New()
	utilsValidations.UseJsonTagAsFieldName(utilsValidations.MyValidation)
	customValidations.RegisterCustomValidations(utilsValidations.MyValidation)

	// Set Routers
	pkgRouters.SetRouter(app)

	// Listen to port
	port := utils.GetEnvValue("APP_PORT", utilsEnums.DefaultAppPort)
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Error when listen to port :%s: %v", port, err)
		return
	}
}

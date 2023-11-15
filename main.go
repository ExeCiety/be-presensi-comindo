package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ExeCiety/be-presensi-comindo/cmd"
	"github.com/ExeCiety/be-presensi-comindo/db"
	pkgRouters "github.com/ExeCiety/be-presensi-comindo/pkg/routers"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
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

	// Create Fiber app
	app := fiber.New()

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

	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Output: logFile,
	}))

	// Set Routers
	pkgRouters.SetRouter(app)

	// Listen to port
	port := utils.GetEnvValue("APP_PORT", utilsEnums.DefaultAppPort)
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Error when listen to port %s: %v", port, err)
		return
	}
}

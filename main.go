package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ExeCiety/be-presensi-comindo/cmd"
	"github.com/ExeCiety/be-presensi-comindo/db"
	pkgRouters "github.com/ExeCiety/be-presensi-comindo/pkg/routers"
	"github.com/ExeCiety/be-presensi-comindo/utils"
	utilsEnums "github.com/ExeCiety/be-presensi-comindo/utils/enums"

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

	// Set Routers
	pkgRouters.SetRouter(app)

	// Listen to port
	port := utils.GetEnvValue("APP_PORT", utilsEnums.DefaultAppPort)
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Error when listen to port %s: %v", port, err)
		return
	}
}

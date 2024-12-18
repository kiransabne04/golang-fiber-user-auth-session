package main

import (
	"fiber-user-auth-session/config"
	"fiber-user-auth-session/internal"
	"fiber-user-auth-session/routers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initialize app config & database connection
	appConfig, err := config.InitializeAppConfig(".")
	if err != nil {
		log.Fatalf("Failed to initialize application configurations: %v", err)
	}

	defer appConfig.Database.Close()

	//initialize all services with AppConfig
	appServices := internal.NewAppServices(appConfig.Database, appConfig.EnvConfig.JWTsecret)

	// create a  new fiber instance
	app := fiber.New(fiber.Config{EnablePrintRoutes: true})

	routers.SetupRouters(app, appServices)

	// Start server
	if err := app.Listen(":" + appConfig.EnvConfig.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

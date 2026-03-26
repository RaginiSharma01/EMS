package main

import (
	"ems/config"
	"ems/db"
	"ems/handler"
	"ems/route"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	cfg := config.LoadConfig()

	config.ConnectRedis()

	database, err := db.ConnectDb(cfg)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	defer database.Pool.Close()

	app := fiber.New(fiber.Config{
		BodyLimit:         20 * 1024 * 1024,
		StreamRequestBody: true,
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("server running")
	})

	employeeHandler := InitializeEmployeeHandler(database.Pool)
	departmentHandler := InitializeDepartmentHandler(database.Pool)
	assetHandler := InitializeAssetHandler(database.Pool)
	salaryHandler := InitializeSalaryHandler(database.Pool)

	route.SetupEmployeeRoutes(
		app,
		employeeHandler,
		departmentHandler,
		assetHandler,
		salaryHandler,
		handler.NewAuthHandler(employeeHandler.Service),
	)

	log.Fatal(app.Listen(cfg.ServerPort))
}

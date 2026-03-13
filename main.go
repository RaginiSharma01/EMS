package main

import (
	"ems/config"
	"ems/db"
	"ems/handler"
	"ems/repository"
	"ems/route"
	"ems/services"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {

	cfg := config.LoadConfig()

	database, err := db.ConnectDb(cfg)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("server running")
	})

	employeeRepo := repository.NewEmployeeRepository(database.Pool)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	route.SetupEmployeeRoutes(app, employeeHandler)

	log.Fatal(app.Listen(cfg.ServerPort))

	defer database.Pool.Close()
}

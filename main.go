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
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {

	cfg := config.LoadConfig()

	database, err := db.ConnectDb(cfg)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	app := fiber.New(fiber.Config{
		BodyLimit:         20 * 1024 * 1024,
		StreamRequestBody: true,
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("server running")
	})
	app.Use("/uploads", static.New("./uploads"))

	// emp
	employeeRepo := repository.NewEmployeeRepository(database.Pool)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	//dpt
	departmentRepo := repository.NewDepartment(database.Pool)
	departmentService := services.NewDepartmentService(departmentRepo)
	departmentHandler := handler.NewDepartmentHandler(departmentService)

	//asset
	assetRepo := repository.NewAssetRepository(database.Pool)
	assetService := services.NewAssetService(assetRepo)
	assetHandler := handler.NewAssetHandler(assetService)

	//salary category
	salaryRepo := repository.NewSalaryCategoryRepository(database.Pool)
	salaryService := services.NewSalaryCategoryService(salaryRepo)
	salaryHandler := handler.NewSalaryCategoryHandler(salaryService)

	route.SetupEmployeeRoutes(
		app,
		employeeHandler,
		departmentHandler,
		assetHandler,
		salaryHandler,
		handler.NewAuthHandler(employeeService),
	)
	log.Fatal(app.Listen(cfg.ServerPort))
	defer database.Pool.Close()
}

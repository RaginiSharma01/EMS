package route

import (
	"ems/handler"
	"ems/middleware"

	"github.com/gofiber/fiber/v3"
)

func SetupEmployeeRoutes(
	app *fiber.App,
	employeeHandler *handler.EmployeeHandler,
	departmentHandler *handler.DepartmentHandler,
	assetHandler *handler.AssetHandler,
	salaryCategoryHandler *handler.SalaryCategoryHandler,
	authHandler *handler.AuthHandler,
) {
	// public routes
	app.Post("/signup", employeeHandler.CreateEmployee)
	app.Post("/login", authHandler.Login)
	app.Get("/all", employeeHandler.GetAllEmployee)
	app.Get("/pdf", employeeHandler.DownloadEmployeePDF)

	// for middleware
	auth := app.Group("/employees", middleware.AuthMiddleware)
	auth.Post("/assets", assetHandler.CreateAsset)
	auth.Post("/assign-asset", assetHandler.AssignAsset)
	auth.Get("/:id", employeeHandler.GetEmployeeByID)
	auth.Post("/logout", middleware.AuthMiddleware, authHandler.Logout)
	auth.Post("/verify-otp", authHandler.VerifyOTP)

	// departments
	departments := app.Group("/departments", middleware.AuthMiddleware)
	departments.Post("/", departmentHandler.CreateDepartment)
	departments.Get("/all", departmentHandler.GetAllDepartment)

	// assets
	assets := app.Group("/assets", middleware.AuthMiddleware)
	assets.Post("/create", assetHandler.CreateAsset)
	assets.Get("/all", assetHandler.GetAllAssets)
}

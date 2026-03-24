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

	employee := app.Group("/employees")
	//signUp---> public route
	app.Post("/signup", employeeHandler.CreateEmployee)
	//login ---> public route
	app.Post("/login", authHandler.Login)

	employee.Use(middleware.AuthMiddleware)
	employee.Get("/all", employeeHandler.GetAllEmployee)
	employee.Get("/:id", employeeHandler.GetEmployeeByID)
	employee.Post("/assign-asset", assetHandler.AssignAsset)
	//logout
	employee.Post("/logout", middleware.AuthMiddleware, authHandler.Logout)

	// departments
	departments := app.Group("/departments", middleware.AuthMiddleware)
	departments.Post("/", departmentHandler.CreateDepartment)
	departments.Get("/all", departmentHandler.GetAllDepartment)

	// assets
	assets := app.Group("/assets", middleware.AuthMiddleware)
	assets.Post("/create", assetHandler.CreateAsset)
	assets.Get("/all", assetHandler.GetAllAssets)
}

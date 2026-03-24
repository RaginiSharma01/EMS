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

	employee.Post("/signup", employeeHandler.CreateEmployee)
	employee.Get("/all", employeeHandler.GetAllEmployee)
	employee.Get("/:id", employeeHandler.GetEmployeeByID)
	employee.Post("/assign-asset", assetHandler.AssignAsset)
	employee.Post("/login", authHandler.Login)

	// logout
	app.Post("/logout",
		middleware.AuthMiddleware,
		authHandler.Logout,
	)

	// departments
	departments := app.Group("/departments")
	departments.Post("/", departmentHandler.CreateDepartment)
	departments.Get("/all", departmentHandler.GetAllDepartment)

	// assets
	assets := app.Group("/assets")
	assets.Post("/create", assetHandler.CreateAsset)
	assets.Get("/all", assetHandler.GetAllAssets)
}

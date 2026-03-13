package route

import (
	"ems/handler"

	"github.com/gofiber/fiber/v3"
)

func SetupEmployeeRoutes(app *fiber.App, employeeHandler *handler.EmployeeHandler, departmentHandler *handler.DepartmentHandler, assetHandler *handler.AssetHandler) {

	employee := app.Group("/employees")

	employee.Post("/", employeeHandler.CreateEmployee)
	employee.Get("/all", employeeHandler.GetAllEmployee)

	departments := app.Group("/departments")
	departments.Post("/", departmentHandler.CreateDepartment)
	departments.Get("/all", departmentHandler.GetAllDepartment)

	assets := app.Group("/assets")
	assets.Post("/", assetHandler.CreateAsset)
}

package route

import (
	"ems/handler"

	"github.com/gofiber/fiber/v3"
)

func SetupEmployeeRoutes(
	app *fiber.App,
	employeeHandler *handler.EmployeeHandler,
	departmentHandler *handler.DepartmentHandler,
	assetHandler *handler.AssetHandler,
	salaryCategoryHandler *handler.SalaryCategoryHandler,
) {

	employee := app.Group("/employees")

	employee.Post("/add-employee", employeeHandler.CreateEmployee)
	employee.Get("/all", employeeHandler.GetAllEmployee)
	employee.Get("/:id", employeeHandler.GetEmployeeByID)
	employee.Post("/:id/assets", assetHandler.AssignAssetToEmployee)

	departments := app.Group("/departments")
	departments.Post("/", departmentHandler.CreateDepartment)
	departments.Get("/all", departmentHandler.GetAllDepartment)

	assets := app.Group("/assets")

	assets.Post("/", assetHandler.CreateAsset)
	assets.Get("/all", assetHandler.GetAllAssets)
	salary := app.Group("/salary-category")
	salary.Post("/", salaryCategoryHandler.CreateCategory)
	salary.Get("/all", salaryCategoryHandler.GetAllCategory)
}

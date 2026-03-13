package route

import (
	"ems/handler"

	"github.com/gofiber/fiber/v3"
)

func SetupEmployeeRoutes(app *fiber.App, employeeHandler *handler.EmployeeHandler) {

	employee := app.Group("/employees")

	employee.Post("/", employeeHandler.CreateEmployee)

}

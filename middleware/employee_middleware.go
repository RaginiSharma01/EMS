package middleware

import (
	"ems/models"

	"github.com/gofiber/fiber/v3"
)

func ValidateEmployee(c fiber.Ctx) error {

	var employee models.Employee

	if err := c.Bind().Body(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if employee.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	if employee.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	if employee.DepartmentID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Department is required",
		})
	}

	return c.Next()
}

package handler

import (
	"ems/models"
	"ems/services"

	"github.com/gofiber/fiber/v3"
)

type EmployeeHandler struct {
	Service *services.EmployeeService
}

func NewEmployeeHandler(service *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		Service: service,
	}
}

// fiber logic for http routing layer

func (h *EmployeeHandler) CreateEmployee(c fiber.Ctx) error {

	var emp models.Employee

	if err := c.Bind().Body(&emp); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	id, err := h.Service.CreateEmployee(c.Context(), emp)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}
package handler

import (
	"ems/models"
	"ems/services"

	"github.com/gofiber/fiber/v3"
)

type DepartmentHandler struct {
	Service *services.DepartmentService
}

func NewDepartmentHandler(service *services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		Service: service,
	}
}

func (h *DepartmentHandler) CreateDepartment(c fiber.Ctx) error {

	var dept models.Department

	err := c.Bind().Body(&dept)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	id, err := h.Service.CreateDepartment(c.Context(), dept)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"id": id,
	})
}

func (h *DepartmentHandler) GetAllDepartment(c fiber.Ctx) error {

	departments, err := h.Service.GetAllDepartment(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(departments)
}


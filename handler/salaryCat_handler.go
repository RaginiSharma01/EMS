package handler

import (
	"ems/models"
	"ems/services"

	"github.com/gofiber/fiber/v3"
)

type SalaryCategoryHandler struct {
	Service *services.SalaryCategoryService
}

func NewSalaryCategoryHandler(service *services.SalaryCategoryService) *SalaryCategoryHandler {
	return &SalaryCategoryHandler{
		Service: service,
	}
}

func (h *SalaryCategoryHandler) CreateCategory(c fiber.Ctx) error {

	var cat models.SalaryCategory

	if err := c.Bind().Body(&cat); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	id, err := h.Service.CreateCategory(c.Context(), cat)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"id": id,
	})
}

func (h *SalaryCategoryHandler) GetAllCategory(c fiber.Ctx) error {

	categories, err := h.Service.GetAllCategory(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(categories)
}

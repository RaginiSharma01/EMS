package handler

import (
	"ems/models"
	"ems/services"

	"github.com/gofiber/fiber/v3"
)

type AssetHandler struct {
	Service *services.AssetService
}

func NewAssetHandler(service *services.AssetService) *AssetHandler {
	return &AssetHandler{
		Service: service,
	}
}

func (h *AssetHandler) CreateAsset(c fiber.Ctx) error {

	var asset models.Asset

	err := c.Bind().Body(&asset)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	id, err := h.Service.CreateAsset(c.Context(), asset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"id": id,
	})
}

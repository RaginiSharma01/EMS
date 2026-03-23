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
	return &AssetHandler{Service: service}
}

func (h *AssetHandler) CreateAsset(c fiber.Ctx) error {

	var asset models.Asset

	err := c.Bind().Body(&asset)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	id, err := h.Service.CreateAsset(c.Context(), asset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"assetId": id,
	})
}

func (h *AssetHandler) GetAllAssets(c fiber.Ctx) error {

	assets, err := h.Service.GetAllAssets(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(assets)
}

func (h *AssetHandler) AssignAssetToEmployee(c fiber.Ctx) error {

	empID := c.Params("id")

	var asset models.Asset

	err := c.Bind().Body(&asset)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	assetID, err := h.Service.CreateAndAssignAsset(
		c.Context(),
		empID,
		asset,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "asset assigned successfully",
		"assetId": assetID,
	})
}

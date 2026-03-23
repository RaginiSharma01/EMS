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

type AssignAssetRequest struct {
	EmpID   string `json:"empId"`
	AssetID string `json:"assetId"`
}

func (h *AssetHandler) AssignAsset(c fiber.Ctx) error {

	var req AssignAssetRequest

	err := c.Bind().Body(&req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err = h.Service.AssignAssetToEmployee(
		c.Context(),
		req.EmpID,
		req.AssetID,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "asset assigned successfully",
	})
}

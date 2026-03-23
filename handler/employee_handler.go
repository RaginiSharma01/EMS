package handler

import (
	"ems/models"
	"ems/services"

	"github.com/gofiber/fiber/v3"
)

type EmployeeHandler struct {
	Service *services.EmployeeService
}
type AuthHandler struct {
	Service *services.EmployeeService
}

func NewEmployeeHandler(service *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		Service: service,
	}
}

func (h *EmployeeHandler) CreateEmployee(c fiber.Ctx) error {

	var req models.CreateEmployeeRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request body",
			"data":    fiber.Map{},
		})
	}

	id, err := h.Service.CreateEmployee(c.Context(), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
			"data":    fiber.Map{},
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Employee added successfully",
		"data": fiber.Map{
			"empId": id,
		},
	})
}

func (h *EmployeeHandler) GetAllEmployee(c fiber.Ctx) error {

	employees, err := h.Service.GetAllEmployee(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
			"data":    fiber.Map{},
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Employees fetched successfully",
		"data":    employees,
	})
}

func (h *EmployeeHandler) GetEmployeeByID(c fiber.Ctx) error {

	id := c.Params("id")

	emp, err := h.Service.GetEmployeeByID(c.Context(), id)
	if err != nil {

		if err.Error() == "no such user exists" {
			return c.Status(400).JSON(fiber.Map{
				"status":  400,
				"message": "No such user exists",
				"data":    fiber.Map{},
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
			"data":    fiber.Map{},
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Employee details fetched successfully",
		"data":    emp,
	})
}
func NewAuthHandler(service *services.EmployeeService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {

	var req models.LoginRequest

	// parse request body
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// call service
	token, err := h.Service.Login(c.Context(), req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return token
	return c.JSON(fiber.Map{
		"token": token,
	})
}

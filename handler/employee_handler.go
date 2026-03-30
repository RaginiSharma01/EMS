package handler

import (
	"bytes"
	"ems/models"
	"ems/services"
	"strconv"
	"strings"

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
		return c.Status(200).JSON(fiber.Map{
			"message": "Invalid request body",
			"data":    fiber.Map{},
		})
	}

	id, err := h.Service.CreateEmployee(c.Context(), req)
	if err != nil {
		return c.Status(200).JSON(fiber.Map{
			"message": err.Error(),
			"data":    fiber.Map{},
		})
	}

	return c.JSON(fiber.Map{
		"message": "Employee added successfully",
		"data": fiber.Map{
			"empId": id,
		},
	})
}

func (h *EmployeeHandler) GetAllEmployee(c fiber.Ctx) error {

	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	employees, err := h.Service.GetAllEmployee(c.Context(), limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"page":  page,
		"limit": limit,
		"data":  employees,
	})
}

func (h *EmployeeHandler) GetEmployeeByID(c fiber.Ctx) error {

	requestedID := c.Params("id")
	tokenEmpID := c.Locals("empId").(string)

	// prevent access to other employees
	if requestedID != tokenEmpID {
		return c.Status(200).JSON(fiber.Map{

			"message": "you cannot access other employee data",
			// "data":fiber.Map{}
		})
	}

	emp, err := h.Service.GetEmployeeByID(c.Context(), requestedID)
	if err != nil {

		if err.Error() == "no such user exists" {
			return c.Status(500).JSON(fiber.Map{

				"message": "No such user exists",
				"data":    fiber.Map{},
			})
		}

		return c.Status(200).JSON(fiber.Map{

			"message": err.Error(),
			"data":    fiber.Map{},
		})
	}

	return c.JSON(fiber.Map{
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
func (h *AuthHandler) Logout(c fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(200).JSON(fiber.Map{
			"message": "token missing",
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer")

	err := h.Service.Logout(c.Context(), token)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "logout successful",
	})
}

func (h *EmployeeHandler) DownloadEmployeePDF(c fiber.Ctx) error {

	// 1. Fetch employees
	employees, err := h.Service.GetAllEmployee(c.Context(), 1000, 0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch employees",
		})
	}

	// 2. Generate PDF
	pdf, err := h.Service.GeneratePdf(employees)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate PDF",
		})
	}

	// 3. Write PDF into a buffer
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to write PDF",
		})
	}

	// 4. Stream buffer to response
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=employees.pdf")
	return c.Send(buf.Bytes())
}

func (h *EmployeeHandler) VerifyEmail(c fiber.Ctx) error {

	token := c.Query("token")

	err := h.Service.EmailVerification(c.Context(), token)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.SendString("Email verified successfully")
}

func (h *EmployeeHandler) VerifyOTP(c fiber.Ctx) error {

	var req models.VerifyOTPRequest

	// parse request body
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	err := h.Service.VerifyOTP(
		c.Context(),
		req.Email,
		req.OTP,
	)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Email verified successfully",
	})
}

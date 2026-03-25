package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"time"

	"ems/config"
	"ems/models"
	"ems/repository"
	"ems/utils"

	"github.com/go-pdf/fpdf"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

func (s *EmployeeService) GetAllEmployee(ctx context.Context, limit int, offset int) ([]models.Employee, error) {
	return s.Repo.GetAllEmployee(ctx, limit, offset)
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id string) (models.EmployeeDetail, error) {
	if id == "" {
		return models.EmployeeDetail{}, errors.New("employee id is required")
	}

	emp, err := s.Repo.GetEmployeeByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.EmployeeDetail{}, errors.New("no such user exists")
		}
		return models.EmployeeDetail{}, err
	}

	return emp, nil
}

func (s *EmployeeService) CreateEmployee(
	ctx context.Context,
	req models.CreateEmployeeRequest,
) (string, error) {

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	fmt.Println("Hashed Password:", hashedPassword)

	// Email validation
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, req.Email)
	if !matched {
		return "", errors.New("incorrect email format")
	}

	// Phone validation
	phoneRegex := `^[0-9]{10}$`
	matched, _ = regexp.MatchString(phoneRegex, req.PhoneNumber)
	if !matched {
		return "", errors.New("incorrect phone number format")
	}

	// JoiningDate validation
	if req.JoiningDate == "" {
		return "", errors.New("joining date is required")
	}

	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, req.JoiningDate)
	if err != nil {
		return "", err
	}

	// Get department ID
	deptID, err := s.Repo.GetDepartmentByName(ctx, req.Department)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", errors.New("no such department exists")
		}
		return "", err
	}

	emp := models.Employee{
		Name:         req.Name,
		Email:        req.Email,
		Password:     hashedPassword,
		PhoneNumber:  req.PhoneNumber,
		DepartmentID: deptID,
		Salary:       req.Salary,
		Location:     req.Location,
		JoiningDate:  parsedTime,
	}

	return s.Repo.CreateEmployee(ctx, emp)
}
func (s *EmployeeService) Login(
	ctx context.Context,
	req models.LoginRequest,
) (string, error) {

	key := "employee:email:" + req.Email
	var emp models.Employee

	val, err := config.RedisClient.Get(ctx, key).Result()

	if err == nil {

		// Redis HIT
		err := json.Unmarshal([]byte(val), &emp)
		if err != nil {
			return "", err
		}

	} else if err == redis.Nil {

		// Redis MISS → check DB
		emp, err = s.Repo.GetEmployeeByEmail(ctx, req.Email)
		if err != nil {
			if err == pgx.ErrNoRows {
				return "", errors.New("Not a registered email please signup!")
			}
			return "", err
		}

		// store in Redis
		userJSON, err := json.Marshal(emp)
		if err == nil {
			//runs the func in seperate light weight treAD
			go config.RedisClient.Set(ctx, key, userJSON, time.Hour*24)
		}

	} else {
		return "", err
	}

	// Verify password
	err = utils.CheckPasswordHash(req.Password, emp.Password)
	if err != nil {
		return "", errors.New("Enter the password again")
	}

	// Generate JWT
	token, err := utils.GenerateJWT(emp.ID, emp.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *EmployeeService) Logout(ctx context.Context, token string) error {
	if token == "" {
		return errors.New("token missing")
	}

	//redis client
	err := config.RedisClient.Set(ctx, token, "blacklisted", time.Hour*24).Err()

	if err != nil {
		return err
	}
	return nil
}

func (s *EmployeeService) GeneratePdf(employees []models.Employee) (*fpdf.Fpdf, error) {

	pdf := fpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	// Title
	pdf.SetFont("Arial", "B", 18)
	pdf.SetTextColor(33, 97, 140)
	pdf.Cell(0, 12, "Employee List")
	pdf.Ln(16)

	headers := []string{"Sl.No", "Name", "Email", "Phone", "Department", "Salary", "Location", "Joining Date"}
	widths := []float64{15, 29, 30, 40, 70, 30, 35, 35}

	// Table header
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(33, 97, 140)
	pdf.SetTextColor(255, 255, 255)
	for i, h := range headers {
		pdf.CellFormat(widths[i], 10, h, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Data rows
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(0, 0, 0)
	for idx, emp := range employees {
		if idx%2 == 0 {
			pdf.SetFillColor(235, 245, 255)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}

		pdf.CellFormat(widths[0], 8, fmt.Sprintf("%d", idx+1), "1", 0, "C", true, 0, "")
		pdf.CellFormat(widths[1], 8, emp.Name, "1", 0, "L", true, 0, "")
		pdf.CellFormat(widths[2], 8, emp.Email, "1", 0, "L", true, 0, "")
		pdf.CellFormat(widths[3], 8, emp.PhoneNumber, "1", 0, "L", true, 0, "")
		pdf.CellFormat(widths[4], 8, emp.DepartmentID, "1", 0, "L", true, 0, "")
		pdf.CellFormat(widths[5], 8, fmt.Sprintf("Rs-%.2f", emp.Salary), "1", 0, "R", true, 0, "")
		pdf.CellFormat(widths[6], 8, emp.Location, "1", 0, "L", true, 0, "")
		pdf.CellFormat(widths[7], 8, emp.JoiningDate.Format("02-Jan-2006"), "1", 1, "C", true, 0, "")
	}

	// Footer
	pdf.Ln(6)
	pdf.SetFont("Arial", "I", 9)
	pdf.SetTextColor(150, 150, 150)
	pdf.Cell(0, 8, fmt.Sprintf("Total Employees: %d", len(employees)))

	return pdf, nil
}

package services

import (
	"context"
	"errors"
	"regexp"
	"time"

	"ems/models"
	"ems/repository"

	"github.com/jackc/pgx/v5"
)

type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		Repo: repo,
	}
}

func (s *EmployeeService) GetAllEmployee(ctx context.Context) ([]models.Employee, error) {
	return s.Repo.GetAllEmployee(ctx)
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
		//fmt.Println("Error:", err)
		return "", err
	}

	// // JoiningDate should not be a future date (optional but good practice)
	// if req.JoiningDate.After(time.Now()) {
	// 	return "", errors.New("joining date cannot be a future date")
	// }

	// Get department ID from name
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
		PhoneNumber:  req.PhoneNumber,
		DepartmentID: deptID,
		Salary:       req.Salary,
		Location:     req.Location,
		JoiningDate:  parsedTime,
	}

	return s.Repo.CreateEmployee(ctx, emp)
}

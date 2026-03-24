package services

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"ems/config"
	"ems/models"
	"ems/repository"
	"ems/utils"

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

	emp, err := s.Repo.GetEmployeeByEmail(ctx, req.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	err = utils.CheckPasswordHash(req.Password, emp.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}

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

	err := config.RedisClient.Set(
		ctx, token, "blacklisted", time.Hour*24,
	).Err()

	if err != nil {
		return err
	}

	return nil
}

package services

import (
	"context"
	"ems/models"
	"ems/repository"
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

func (s *EmployeeService) CreateEmployee(
	ctx context.Context,
	emp models.Employee,
) (string, error) {

	return s.Repo.CreateEmployee(ctx, emp)

}

func (s *EmployeeService) UpdateProfileImage(ctx context.Context, id string, filename string) error {
	return s.Repo.UpdateProfileImage(ctx, id, filename)
}

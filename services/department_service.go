package services

import (
	"context"
	"ems/models"
	"ems/repository"
)

type DepartmentService struct {
	Repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		Repo: repo,
	}
}

func (s *DepartmentService) CreateDepartment(
	ctx context.Context,
	dept models.Department,
) (string, error) {

	return s.Repo.CreateDepartment(ctx, dept)
}

func (s *DepartmentService) GetAllDepartment(ctx context.Context) ([]models.Department, error) {
	return s.Repo.GetAllDepartment(ctx)
}

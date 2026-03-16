package services

import (
	"context"
	"ems/models"
	"ems/repository"
)

type SalaryCategoryService struct {
	Repo *repository.SalaryCategoryRepository
}

func NewSalaryCategoryService(repo *repository.SalaryCategoryRepository) *SalaryCategoryService {
	return &SalaryCategoryService{
		Repo: repo,
	}
}

func (s *SalaryCategoryService) CreateCategory(ctx context.Context, cat models.SalaryCategory) (string, error) {

	return s.Repo.CreateCategory(ctx, cat)

}

func (s *SalaryCategoryService) GetAllCategory(ctx context.Context) ([]models.SalaryCategory, error) {

	return s.Repo.GetAllCategory(ctx)

}

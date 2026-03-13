package services

import (
	"context"
	"ems/models"
	"ems/repository"
)

type AssetService struct {
	Repo *repository.AssetsRepository
}

func NewAssetService(repo *repository.AssetsRepository) *AssetService {
	return &AssetService{
		Repo: repo,
	}
}

func (s *AssetService) CreateAsset(
	ctx context.Context,
	asset models.Asset,
) (string, error) {

	return s.Repo.CreateAsset(ctx, asset)
}

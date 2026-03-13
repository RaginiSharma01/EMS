package repository

import (
	"context"
	"ems/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

//assests and employee assets query goes here

type AssetsRepository struct {
	DB *pgxpool.Pool
}

func NewAssetRepository(pool *pgxpool.Pool) *AssetsRepository {
	return &AssetsRepository{
		DB: pool,
	}
}

func (r *AssetsRepository) CreateAsset(
	ctx context.Context,
	asset models.Asset,
) (string, error) {

	query := `
	INSERT INTO assets
	(asset_name, asset_type, asset_price, dept_id)
	VALUES ($1, $2, $3, $4)
	RETURNING asset_id
	`

	var id string

	err := r.DB.QueryRow(
		ctx,
		query,
		asset.AssetName,
		asset.AssetType,
		asset.AssetPrice,
		asset.DeptID,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

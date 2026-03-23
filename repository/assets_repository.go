package repository

import (
	"context"
	"ems/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AssetsRepository struct {
	DB *pgxpool.Pool
}

func NewAssetRepository(pool *pgxpool.Pool) *AssetsRepository {
	return &AssetsRepository{DB: pool}
}

func (r *AssetsRepository) CreateAsset(
	ctx context.Context,
	asset models.Asset,
) (string, error) {

	query := `
	INSERT INTO assets
	(asset_name, asset_type, asset_price, dept_id)
	VALUES ($1,$2,$3,$4)
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

func (r *AssetsRepository) GetDepartmentIDByName(
	ctx context.Context,
	name string,
) (string, error) {

	query := `
SELECT dept_id
FROM departments
WHERE name = $1
`

	var deptID string

	err := r.DB.QueryRow(ctx, query, name).Scan(&deptID)
	if err != nil {
		return "", err
	}

	return deptID, nil
}

func (r *AssetsRepository) GetAllAssets(ctx context.Context) ([]models.Asset, error) {

	query := `
	SELECT asset_id, asset_name, asset_type, asset_price, dept_id
	FROM assets
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []models.Asset

	for rows.Next() {

		var asset models.Asset

		err := rows.Scan(
			&asset.AssetID,
			&asset.AssetName,
			&asset.AssetType,
			&asset.AssetPrice,
			&asset.DeptID,
		)

		if err != nil {
			return nil, err
		}

		assets = append(assets, asset)
	}

	return assets, nil
}

func (r *AssetsRepository) AssignAssetToEmployee(
	ctx context.Context,
	empID string,
	assetID string,
) error {

	query := `
	INSERT INTO emp_asset
	(emp_id, asset_id)
	VALUES ($1,$2)
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		empID,
		assetID,
	)

	return err
}

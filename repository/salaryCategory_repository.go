package repository

//salary category query goes here .
// salary of category of a employee goes here .



import (
	"context"
	"ems/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SalaryCategoryRepository struct {
	DB *pgxpool.Pool
}

func NewSalaryCategoryRepository(pool *pgxpool.Pool) *SalaryCategoryRepository {
	return &SalaryCategoryRepository{
		DB: pool,
	}
}

func (r *SalaryCategoryRepository) CreateCategory(
	ctx context.Context,
	cat models.SalaryCategory,
) (string, error) {

	query := `
	INSERT INTO salary_category (min_sal, max_sal, cat_name)
	VALUES ($1,$2,$3)
	RETURNING cat_id
	`

	var id string

	err := r.DB.QueryRow(
		ctx,
		query,
		cat.MinSal,
		cat.MaxSal,
		cat.CatName,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *SalaryCategoryRepository) GetAllCategory(ctx context.Context) ([]models.SalaryCategory, error) {

	query := `
	SELECT cat_id, min_sal, max_sal, cat_name
	FROM salary_category
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.SalaryCategory

	for rows.Next() {

		var cat models.SalaryCategory

		err := rows.Scan(
			&cat.CatID,
			&cat.MinSal,
			&cat.MaxSal,
			&cat.CatName,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, cat)
	}

	return categories, nil
}
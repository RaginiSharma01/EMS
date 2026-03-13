package repository

import (
	"context"
	"ems/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DepartmentRepository struct {
	DB *pgxpool.Pool
}

func NewDepartment(pool *pgxpool.Pool) *DepartmentRepository {
	return &DepartmentRepository{
		DB: pool,
	}
}

func (r *DepartmentRepository) CreateDepartment(
	ctx context.Context,
	dept models.Department,
) (string, error) {

	query := `
	INSERT INTO departments (name, location)
	VALUES ($1, $2)
	RETURNING dept_id
	`

	var id string

	err := r.DB.QueryRow(
		ctx,
		query,
		dept.Name,
		dept.Location,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
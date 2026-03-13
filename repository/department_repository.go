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

func (r *DepartmentRepository) GetAllDepartment(ctx context.Context) ([]models.Department, error) {
	query := `SELECT dept_id, name, location, created_at, updated_at, FROM departments ORDER BY created_at DESC`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var departments []models.Department

	for rows.Next() {
		var dept models.Department

		err := rows.Scan(
			&dept.DeptID,
			&dept.Name,
			&dept.Location,
			&dept.CreatedAt,
			&dept.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		departments = append(departments, dept)
	}
	return departments, nil

}

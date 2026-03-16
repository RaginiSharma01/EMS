package repository

import (
	"context"
	"ems/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct {
	DB *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{
		DB: pool,
	}
}

func (r *EmployeeRepository) GetAllEmployee(ctx context.Context) ([]models.Employee, error) {

	query := `
	SELECT 
	e.id,
	e.name,
	e.email,
	e.department_id,
	e.salary,
	e.location,
	e.joining_date,
	e.created_at,
	e.updated_at
	FROM employees_data e
	LEFT JOIN departments d
	ON e.department_id = d.dept_id
	ORDER BY e.created_at DESC
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {

		var emp models.Employee

		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Email,
			&emp.DepartmentID,
			&emp.Salary,
			&emp.Location,
			&emp.JoiningDate,
			&emp.CreatedAt,
			&emp.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

func (r *EmployeeRepository) CreateEmployee(
	ctx context.Context,
	emp models.Employee,
) (string, error) {

	query := `
	INSERT INTO employees_data
	(name, email, department_id, salary, location, joining_date)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id
	`

	var id string

	err := r.DB.QueryRow(
		ctx,
		query,
		emp.Name,
		emp.Email,
		emp.DepartmentID,
		emp.Salary,
		emp.Location,
		emp.JoiningDate,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

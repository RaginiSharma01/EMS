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
id,
name,
email,
department_id, //name
salary,
location,
joining_date,
created_at,
updated_at,
profile_image
FROM employees_data
ORDER BY created_at DESC //assets , 
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
			&emp.ProfileImage,
			&emp.ProfileImage,
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

func (r *EmployeeRepository) UpdateProfileImage(ctx context.Context, id string, filename string) error {
	query := `UPDATE employees_data SET profile_image = $1 WHERE id = $2`

	_, err := r.DB.Exec(ctx, query, filename, id)
	return err
}

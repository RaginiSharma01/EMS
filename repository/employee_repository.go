package repository

import (
	"context"
	"ems/models"

	"github.com/jackc/pgx/v5"
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

func (r *EmployeeRepository) GetAllEmployee(ctx context.Context, limit int, offset int) ([]models.Employee, error) {

	query := `
	SELECT 
		id,
		name,
		email,
		phone_number,
		department_id,
		salary,
		location,
		joining_date,
		created_at,
		updated_at
	FROM employees_data
	LIMIT $1 OFFSET $2
	`

	rows, err := r.DB.Query(ctx, query, limit, offset)
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
			&emp.PhoneNumber,
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *EmployeeRepository) GetEmployeeByID(ctx context.Context, id string) (models.EmployeeDetail, error) {

	// Step 1: Fetch employee + department + salary_category
	// salary_category is computed using CASE in SQL
	empQuery := `
	SELECT 
		e.id,
		e.name,
		e.email,
		e.phone_number,
		e.salary,
		CASE
			WHEN e.salary < 30000                    THEN 'Junior Level'
			WHEN e.salary >= 30000 AND e.salary < 60000 THEN 'Mid Level'
			WHEN e.salary >= 60000 AND e.salary < 100000 THEN 'Senior Level'
			ELSE 'Executive Level'
		END AS salary_category,
		e.location,
		e.joining_date,
		d.name     AS dept_name,
		d.location AS dept_location
	FROM employees_data e
	JOIN departments d ON e.department_id = d.dept_id
	WHERE e.id = $1
	`

	var emp models.EmployeeDetail

	err := r.DB.QueryRow(ctx, empQuery, id).Scan(
		&emp.ID,
		&emp.Name,
		&emp.Email,
		&emp.PhoneNumber,
		&emp.Salary,
		&emp.SalaryCategory,
		&emp.Location,
		&emp.JoiningDate,
		&emp.Department.Name,
		&emp.Department.Location,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.EmployeeDetail{}, pgx.ErrNoRows
		}
		return models.EmployeeDetail{}, err
	}

	// Step 2: Fetch assets assigned to this employee via emp_asset junction table
	assetsQuery := `
	SELECT 
		a.asset_name,
		a.asset_type,
		a.asset_price
	FROM emp_asset ea
	JOIN assets a ON ea.asset_id = a.asset_id
	WHERE ea.emp_id = $1
	`

	rows, err := r.DB.Query(ctx, assetsQuery, id)
	if err != nil {
		return models.EmployeeDetail{}, err
	}
	defer rows.Close()

	// Initialize as empty slice so JSON returns [] not null
	assets := []models.AssetSummary{}

	for rows.Next() {
		var asset models.AssetSummary

		err := rows.Scan(
			&asset.AssetName,
			&asset.AssetType,
			&asset.AssetPrice,
		)
		if err != nil {
			return models.EmployeeDetail{}, err
		}

		assets = append(assets, asset)
	}

	if err := rows.Err(); err != nil {
		return models.EmployeeDetail{}, err
	}

	emp.Assets = assets

	return emp, nil
}

func (r *EmployeeRepository) CreateEmployee(
	ctx context.Context,
	emp models.Employee,
) (string, error) {

	query := `
	INSERT INTO employees_data
		(name, email, password, phone_number, department_id, salary, location, joining_date)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id
	`

	var id string

	err := r.DB.QueryRow(
		ctx,
		query,
		emp.Name,
		emp.Email,
		emp.Password,
		emp.PhoneNumber,
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

func (r *EmployeeRepository) GetDepartmentByName(ctx context.Context, name string) (string, error) {
	query := `SELECT dept_id FROM departments WHERE name=$1`

	var id string
	err := r.DB.QueryRow(ctx, query, name).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *EmployeeRepository) GetEmployeeByEmail(
	ctx context.Context,
	email string,
) (models.Employee, error) {

	query := `
	SELECT id, email, password
	FROM employees_data
	WHERE email=$1
	`

	var emp models.Employee

	err := r.DB.QueryRow(ctx, query, email).Scan(&emp.ID, &emp.Email, &emp.Password)

	return emp, err
}

func (r *EmployeeRepository) MarkEmailVerified(
	ctx context.Context,
	email string,
) error {

	query := `
	UPDATE employees_data
	SET email_verified = true
	WHERE email = $1
	`

	_, err := r.DB.Exec(ctx, query, email)
	return err
}

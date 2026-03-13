package repository

import (
	"context"
	"ems/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct {
	DB *pgxpool.Pool
}


func NewEmployeeRepository (pool *pgxpool.Pool) *EmployeeRepository{
	return &EmployeeRepository{
		DB:pool,
	}
}

func(r *EmployeeRepository) CreateEmployee(ctx context.Context,emp models.Employee) (string , error){
	query := `
	INSERT INTO employees_data
	(name , email , department_id,salary , location , joining_date)
	values($1 , $2 , $3, $4 , $5 ,$6)
	RETURNING id
	`

	var id string

	err := r.DB.QueryRow(ctx,query,
	emp.Name,
	emp.Email,
	emp.DepartmentID,
	emp.Salary,
	emp.Location,
	emp.JoiningDate,
	).Scan(&id)

	if err !=nil{
		return "" , err
	}
	return id , nil

}

func (r *EmployeeRepository) GetAllEmployee(ctx context.Context,limit int,offset int,)([]models.Employee, error) {

	query := `SELECT id, name, email, department_id, salary, location, joining_date
	FROM employees_data
	ORDER BY created_at DESC
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
			&emp.DepartmentID,
			&emp.Salary,
			&emp.Location,
			&emp.JoiningDate,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}
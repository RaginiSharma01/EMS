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
	(name , email , department_id,salary , location , joining_data)
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

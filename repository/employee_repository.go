package repository

import "github.com/jackc/pgx/v5/pgxpool"

type EmployeeRepository struct {
	DB *pgxpool.Pool
}


func employee(pool *pgxpool.Pool) *EmployeeRepository{
	return &EmployeeRepository{
		DB: pool,
	}

	
}
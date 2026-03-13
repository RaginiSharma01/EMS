package models

import "time"

type Employee struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	DepartmentID string    `json:"department_id"`
	Salary       float64   `json:"salary"`
	Location     string    `json:"location"`
	JoiningDate  time.Time `json:"joining_date"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

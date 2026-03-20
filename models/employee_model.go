package models

import "time"

type Employee struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phoneNumber"`
	DepartmentID string    `json:"departmentId"`
	Salary       float64   `json:"salary"`
	Location     string    `json:"location"`
	JoiningDate  time.Time `json:"joiningDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Department struct {
	DeptID   string `json:"dept_id,omitempty"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type AssetSummary struct {
	AssetName  string  `json:"asset_name"`
	AssetType  string  `json:"asset_type"`
	AssetPrice float64 `json:"asset_price"`
}

type EmployeeDetail struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	PhoneNumber    string         `json:"phoneNumber"`
	Salary         float64        `json:"salary"`
	SalaryCategory string         `json:"salary_category"`
	Location       string         `json:"location"`
	JoiningDate    time.Time      `json:"joining_date"`
	Department     Department     `json:"department"`
	Assets         []AssetSummary `json:"assets"`
}

type CreateEmployeeRequest struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phoneNumber"`
	Department  string  `json:"department"`
	Salary      float64 `json:"salary"`
	Location    string  `json:"location"`
	JoiningDate string  `json:"joiningDate"`
}

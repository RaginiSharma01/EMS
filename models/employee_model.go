package models

type Employee struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	DepartmentID string  `json:"department_id"`
	Salary       float64 `json:"salary"`
	Location     string  `json:"location"`
	JoiningDate  string  `json:"joining_date,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}
type Department struct {
	DeptID    string `json:"dept_id"`
	Name      string `json:"name"`
	Location  string `json:"location,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
type Asset struct {
	AssetID    string  `json:"asset_id"`
	AssetName  string  `json:"asset_name"`
	AssetType  string  `json:"asset_type"`
	AssetPrice float64 `json:"asset_price"`
	DeptID     string  `json:"dept_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
type EmployeeAsset struct {
	EmpAssetID string `json:"emp_asset_id"`
	EmpID      string `json:"emp_id"`
	AssetID    string `json:"asset_id"`
}

type SalaryCategory struct {
	CatID   string  `json:"cat_id"`
	CatName string  `json:"cat_name"`
	MinSal  float64 `json:"min_sal"`
	MaxSal  float64 `json:"max_sal"`
}

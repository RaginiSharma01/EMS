package models

import "time"

type Asset struct {
	AssetID        string    `json:"assetId,omitempty"`
	AssetName      string    `json:"assetName"`
	AssetType      string    `json:"assetType"`
	AssetPrice     float64   `json:"assetPrice"`
	DeptID         string    `json:"deptId,omitempty"`
	DepartmentName string    `json:"departmentName,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
}
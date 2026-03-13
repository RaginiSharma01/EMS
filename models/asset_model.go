package models

import "time"

type Asset struct {
	AssetID    string    `json:"asset_id"`
	AssetName  string    `json:"asset_name"`
	AssetType  string    `json:"asset_type"`
	AssetPrice float64   `json:"asset_price"`
	DeptID     string    `json:"dept_id"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
package models

type SalaryCategory struct {
	CatID  string  `json:"cat_id"`
	CatName string `json:"cat_name"`
	MinSal float64 `json:"min_sal"`
	MaxSal float64 `json:"max_sal"`
}
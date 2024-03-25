package models

import "time"

type ProductModel struct {
	ID          int        `json:"id_product" db:"id_product"`
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	CategoryID  string     `json:"id_category" db:"id_category"`
	Created_at  *time.Time `json:"created_at,omitempty" db:"created_at"`
	Updated_at  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

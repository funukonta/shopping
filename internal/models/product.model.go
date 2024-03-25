package models

type ProductModel struct {
	ID          int     `json:"id_product" db:"id_product"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryID  int     `json:"id_category" db:"id_category"`
}

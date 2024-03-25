package models

import "time"

type ShoppingCart struct {
	CartID int `json:"id_cart" db:"id_cart"`
	CustID int `json:"id_customer" db:"id_customer"`
}

type ShoppingCartDetail struct {
	CartIDdtl int        `json:"id_cart_detail" db:"id_cart_detail"`
	CartID    int        `json:"id_cart,omitempty" db:"id_cart"`
	ProductID int        `json:"id_product" db:"id_product"`
	Quantity  int        `json:"qty" db:"qty"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type ResJoinCart struct {
	Cart ShoppingCart         `json:"shopping-cart"`
	Dtl  []ShoppingCartDetail `json:"dtl,omitempty"`
}

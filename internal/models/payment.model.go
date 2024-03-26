package models

import "time"

type PaymentModel struct {
	PaymentID     int       `json:"id_payment" db:"id_payment"`
	InvoiceId     int       `json:"id_invoice" db:"id_invoice"`
	PaymentAmount float64   `json:"payment_amount" db:"payment_amount"`
	PaymentMethod int       `json:"payment_method" db:"payment_method"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type PaymentMethod struct {
	MethodID int    `json:"id_payment_method" db:"id_payment_method"`
	Name     string `json:"name" db:"name"`
}

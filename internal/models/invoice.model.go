package models

import "time"

type Invoice struct {
	InvoiceID  int        `json:"id_invoice" db:"id_invoice"`
	CustomerID int        `json:"id_customer" db:"id_customer"`
	Status     string     `json:"status" db:"status"`
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type InvoiceDetail struct {
	InvoiceDetailID int     `json:"id_invoice_detail" db:"id_invoice_detail"`
	InvoiceID       int     `json:"id_invoice" db:"id_invoice"`
	ProductID       int     `json:"id_product" db:"id_product"`
	Quantity        int     `json:"qty" db:"qty"`
	SubTotal        float64 `json:"sub_total" db:"sub_total"`
}

type BodInvoiceDtl struct {
	ProductID int     `json:"id_product"`
	Quantity  int     `json:"qty"`
	Price     float64 `json:"price"`
}

type ResJoinInvoice struct {
	Inv    Invoice         `json:"inv"`
	InvDtl []InvoiceDetail `json:"dtl"`
}

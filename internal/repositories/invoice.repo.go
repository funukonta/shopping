package repositories

import (
	"context"

	"github.com/funukonta/shopping/internal/models"
	"github.com/jmoiron/sqlx"
)

type InvoiceRepo interface {
	Create(id_cust int, dtl []models.InvoiceDetail) (*models.Invoice, []models.InvoiceDetail, error)
	GetInvoice(id_cust int) ([]models.ResJoinInvoice, error)
	GetInvoiceById(id_cust, id_invoice int) (*models.Invoice, []models.InvoiceDetail, error)
}

type invoiceRepo struct {
	db *sqlx.DB
}

func NewInvoiceRepo(db *sqlx.DB) InvoiceRepo {
	return &invoiceRepo{db: db}
}
func (r *invoiceRepo) Create(id_cust int, dtl []models.InvoiceDetail) (*models.Invoice, []models.InvoiceDetail, error) {

	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	query := `insert into invoice (id_customer) values ($1)
	returning id_invoice, id_customer, status, created_at;`
	inv := new(models.Invoice)
	err = tx.Get(inv, query, id_cust)
	if err != nil {
		return nil, nil, err
	}

	query = `insert into invoice_detail (id_invoice,id_product,qty,sub_total) values ($1,$2,$3,$4)
	returning *`
	details := []models.InvoiceDetail{}
	for _, invdtl := range dtl {
		queryPrice := `select price from products where id_product=$1`
		var price float64
		err := tx.QueryRowx(queryPrice, invdtl.ProductID).Scan(&price)
		if err != nil {
			return nil, nil, err
		}

		invDetail := models.InvoiceDetail{}
		err = tx.Get(&invDetail, query, inv.InvoiceID, invdtl.ProductID, invdtl.Quantity, price*float64(invdtl.Quantity))
		if err != nil {
			return nil, nil, err
		}
		details = append(details, invDetail)
	}
	err = tx.Commit()
	return inv, details, err
}
func (r *invoiceRepo) GetInvoice(id_cust int) ([]models.ResJoinInvoice, error) {
	query := `select id_invoice,status,created_at from invoice where id_customer=$1`
	inv := []models.Invoice{}
	err := r.db.Select(&inv, query, id_cust)
	if err != nil {
		return nil, err
	}

	invoices := []models.ResJoinInvoice{}
	for _, data := range inv {
		query = `select * from invoice_detail where id_invoice=$1`
		invdtl := []models.InvoiceDetail{}
		err = r.db.Select(&invdtl, query, data.InvoiceID)
		if err != nil {
			return nil, err
		}
		invoice := models.ResJoinInvoice{
			Inv:    data,
			InvDtl: invdtl,
		}
		invoices = append(invoices, invoice)
	}

	return invoices, err
}
func (r *invoiceRepo) GetInvoiceById(id_cust, id_invoice int) (*models.Invoice, []models.InvoiceDetail, error) {
	query := `select id_invoice,status,created_at from invoice where id_customer=$1 and id_invoice=$2`
	inv := &models.Invoice{}
	err := r.db.Get(inv, query, id_cust, id_invoice)
	if err != nil {
		return nil, nil, err
	}

	query = `select * from invoice_detail where id_invoice=$1`
	invdtl := []models.InvoiceDetail{}
	err = r.db.Select(&invdtl, query, id_invoice)
	if err != nil {
		return nil, nil, err
	}

	return inv, invdtl, err
}

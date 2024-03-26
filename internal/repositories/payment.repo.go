package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/funukonta/shopping/internal/models"
	"github.com/jmoiron/sqlx"
)

type PaymentRepo interface {
	DoPayment(int, *models.PaymentModel) (*models.PaymentModel, error)
}

type paymentRepo struct {
	db *sqlx.DB
}

func NewPaymentRepo(db *sqlx.DB) PaymentRepo {
	return &paymentRepo{db: db}
}

func (r *paymentRepo) DoPayment(id_invoice int, payment *models.PaymentModel) (*models.PaymentModel, error) {
	// payment method 1:cash 2:digital
	query := `insert into payment (id_invoice,payment_amount,payment_method) values ($1,$2,$3)
	returning *`

	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = tx.Get(payment, query, id_invoice, payment.PaymentAmount, payment.PaymentMethod)
	if err != nil {
		if strings.Contains(err.Error(), "fk_invoice_id") {
			return nil, fmt.Errorf("invoice tidak ditemukan")
		}
		return nil, fmt.Errorf("terjadi kesalahan saat payment")
	}

	query = `select sum(payment_amount) from payment where id_invoice=$1`
	var total_sudah_bayar float64
	err = tx.QueryRowx(query, id_invoice).Scan(&total_sudah_bayar)
	if err != nil {
		return nil, err
	}

	query = `select sum(b.sub_total) from invoice a
	left join invoice_detail b on a.id_invoice = b.id_invoice
	where a.id_invoice=$1;`
	var total_yang_harus_dibayarkan float64
	err = tx.QueryRowx(query, id_invoice).Scan(&total_yang_harus_dibayarkan)
	if err != nil {
		return nil, err
	}

	var status string
	calculate := total_yang_harus_dibayarkan - (total_sudah_bayar)
	if calculate == 0 {
		status = "PAID"
	} else if calculate > 0 {
		status = "BELUM LUNAS"
	} else {
		status = "LEBIH BAYAR"
	}

	query = `update invoice set status=$1 where id_invoice=$2`
	aff, err := tx.Exec(query, status, id_invoice)
	if err != nil {
		return nil, err
	}
	affRow, err := aff.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affRow == 0 {
		return nil, fmt.Errorf("terjadi kesalahan saat payment, update invoice")
	}

	err = tx.Commit()

	return payment, err
}

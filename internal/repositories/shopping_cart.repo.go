package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/funukonta/shopping/internal/models"
	"github.com/jmoiron/sqlx"
)

type ShopCartRepo interface {
	GetAllList(id_cust int) (int, []models.ShoppingCartDetail, error)
	AddCart(int, *models.ShoppingCartDetail) error
	DeleteProduct(int, []models.ShoppingCartDetail) error
	UpdateProduct(*models.ShoppingCart, []models.ShoppingCartDetail) error
}

type shopCartRepo struct {
	db *sqlx.DB
}

func NewShoppingCartRepo(db *sqlx.DB) ShopCartRepo {
	return &shopCartRepo{db: db}
}

func (r *shopCartRepo) GetAllList(id_cust int) (int, []models.ShoppingCartDetail, error) {
	query := `SELECT id_cart from shopping_cart where id_customer=$1`
	var id_cart int
	err := r.db.QueryRowx(query, id_cust).Scan(&id_cart)
	if err != nil {
		return 0, nil, err
	}

	query = `SELECT id_cart_detail, id_product, qty
	From shopping_cart_detail
	WHERE id_cart=$1;`
	cart := []models.ShoppingCartDetail{}
	err = r.db.Select(&cart, query, id_cart)
	return id_cart, cart, err
}

func (r *shopCartRepo) AddCart(id_cust int, product *models.ShoppingCartDetail) error {
	query := `SELECT id_cart from shopping_cart where id_customer=$1`
	var id_cart int
	err := r.db.QueryRowx(query, id_cust).Scan(&id_cart)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if id_cart == 0 {
		query := `insert into shopping_cart (id_customer) values ($1)
		returning id_cart`
		err := tx.QueryRowx(query, id_cust).Scan(&id_cart)
		if err != nil {
			return err
		}
	}

	query = `select id_product from shopping_cart_detail where id_cart=$1 and id_product=$2`
	var id_product_detail int
	err = tx.QueryRowx(query, id_cart, product.ProductID).Scan(&id_product_detail)
	if err != nil && strings.Contains(err.Error(), "no rows") {
		query = `INSERT INTO shopping_cart_detail (id_cart,id_product,qty) values ($1,$2,$3)`
		aff, err := tx.Exec(query, id_cart, product.ProductID, product.Quantity)
		if err != nil {
			return err
		}
		affRow, err := aff.RowsAffected()
		if err != nil {
			return err
		}
		if affRow == 0 {
			return fmt.Errorf("terjadi kesalahan insert shopping cart")
		}
	} else {
		query = `update shopping_cart_detail set qty=qty+$1 where id_cart=$2 and id_product=$3`
		aff, err := tx.Exec(query, product.Quantity, id_cart, product.ProductID)
		if err != nil {
			return err
		}
		affRow, err := aff.RowsAffected()
		if err != nil {
			return err
		}
		if affRow == 0 {
			return fmt.Errorf("terjadi kesalahan insert shopping cart")
		}
	}

	err = tx.Commit()
	return err
}

func (r *shopCartRepo) DeleteProduct(id_cust int, cartDtl []models.ShoppingCartDetail) error {
	query := `select id_cart from shopping_cart where id_customer=$1`
	var id_cart int
	err := r.db.QueryRowx(query, id_cust).Scan(&id_cart)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query = `DELETE from shopping_cart_detail where id_cart=$1 and id_product=$2`
	for _, detail := range cartDtl {
		aff, err := tx.Exec(query, id_cart, detail.ProductID)
		if err != nil {
			return err
		}
		affRow, err := aff.RowsAffected()
		if err != nil {
			return err
		}
		if affRow == 0 {
			return fmt.Errorf("terjadi kesalahan delete shopping cart, cek apakah product ada dalam cart")
		}
	}

	err = tx.Commit()
	return err
}

func (r *shopCartRepo) UpdateProduct(cart *models.ShoppingCart, cartDtl []models.ShoppingCartDetail) error {
	return nil
}

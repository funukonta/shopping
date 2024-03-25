package repositories

import (
	"github.com/funukonta/shopping/internal/models"
	"github.com/jmoiron/sqlx"
)

type ShopCartRepo interface {
	GetAllList(id_cust int) (int, []models.ShoppingCartDetail, error)
	AddCart(*models.ShoppingCart, []models.ShoppingCartDetail) error
	DeleteProduct(*models.ShoppingCart, []models.ShoppingCartDetail) error
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

func (r *shopCartRepo) AddCart(cart *models.ShoppingCart, cartDtl []models.ShoppingCartDetail) error {
	return nil
}

func (r *shopCartRepo) DeleteProduct(cart *models.ShoppingCart, cartDtl []models.ShoppingCartDetail) error {
	return nil
}

func (r *shopCartRepo) UpdateProduct(cart *models.ShoppingCart, cartDtl []models.ShoppingCartDetail) error {
	return nil
}

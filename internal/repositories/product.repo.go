package repositories

import (
	"context"

	"github.com/funukonta/shopping/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	Create(*models.ProductModel) (*models.ProductModel, error)
	GetAll() ([]models.ProductModel, error)
	GetByCategory(string) ([]models.ProductModel, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) Create(data *models.ProductModel) (*models.ProductModel, error) {
	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `INSERT INTO products (name, price, description, id_category)
	VALUES  ($1,$2,$3,$4)
	RETURNING id_product, name, price, description, id_category, created_at;`

	newProduct := new(models.ProductModel)
	if err := tx.Get(newProduct, query, data.Name, data.Price, data.Description, data.CategoryID); err != nil {
		return nil, err
	}
	err = tx.Commit()

	return newProduct, err
}

func (r *productRepo) GetAll() ([]models.ProductModel, error) {
	query := `select id_product, name, price, description, id_category, created_at from products
	order by updated_at desc`
	products := []models.ProductModel{}
	err := r.db.Select(&products, query)
	return products, err
}

func (r *productRepo) GetByCategory(cat string) ([]models.ProductModel, error) {
	query := `select id_product, name, price, description, id_category, created_at from products
	where id_category=$1
	order by updated_at desc`
	products := []models.ProductModel{}
	err := r.db.Select(&products, query, cat)
	return products, err
}

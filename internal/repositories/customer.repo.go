package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/funukonta/shopping/internal/models"
	"github.com/jmoiron/sqlx"
)

type CustomerRepo interface {
	Create(data *models.CustomerModel) (*models.CustomerModel, error)
	GetAll() ([]models.CustomerModel, error)
	GetById(id int) (*models.CustomerModel, error)
	Update(data *models.CustomerModel) error
	Delete(id int) error
}

type customerRepo struct {
	db *sqlx.DB
}

func NewCustomerRepo(db *sqlx.DB) CustomerRepo {
	return &customerRepo{db: db}
}

func (r *customerRepo) Create(data *models.CustomerModel) (*models.CustomerModel, error) {
	query := `INSERT INTO customers (name, email, password, address, phone) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id ,name, email, password, address, phone, created_at
	`
	newCust := new(models.CustomerModel)
	err := r.db.Get(newCust, query, data.Name, data.Email, data.Password, data.Address, data.Phone)
	if err != nil {
		return nil, err
	}

	return newCust, err
}

func (r *customerRepo) GetAll() ([]models.CustomerModel, error) {
	query := `select id,name, email, password, address, phone from customers`
	customers := []models.CustomerModel{}
	if err := r.db.Select(&customers, query); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepo) GetById(id int) (*models.CustomerModel, error) {
	query := `select id, name, email, password, address, phone from customers
	where id=$1`
	customer := new(models.CustomerModel)
	if err := r.db.Get(customer, query, id); err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepo) Update(data *models.CustomerModel) error {
	query := `UPDATE customers 
	SET 
		name = $1,
		email = $2,
		password = $3,
		address = $4,
		phone = $5,
		updated_at = $6
	WHERE
		id = $7;`
	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	aff, err := tx.Exec(query, data.Name, data.Email, data.Password, data.Address, data.Phone, time.Now(), data.ID)
	if err != nil {
		return err
	}
	rowAff, err := aff.RowsAffected()
	if err != nil {
		return err
	}
	if rowAff == 0 {
		return fmt.Errorf("ada kesalahan pada data, cek kembali")
	}

	err = tx.Commit()
	return err
}

func (r *customerRepo) Delete(id int) error {
	return nil
}

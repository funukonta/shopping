package repositories

import (
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
	return nil, nil
}

func (r *customerRepo) GetById(id int) (*models.CustomerModel, error) {
	return nil, nil
}

func (r *customerRepo) Update(data *models.CustomerModel) error {
	return nil
}

func (r *customerRepo) Delete(id int) error {
	return nil
}

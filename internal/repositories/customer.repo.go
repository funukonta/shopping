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
	return nil, nil
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

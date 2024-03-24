package services

import (
	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/repositories"
)

type CustomerService interface {
	Create(data *models.CustomerModel) (*models.CustomerModel, error)
	GetAll() ([]models.CustomerModel, error)
	GetById(id int) (*models.CustomerModel, error)
	Update(data *models.CustomerModel) error
	Delete(id int) error
}

type customerService struct {
	repo repositories.CustomerRepo
}

func NewCustomerService(repo repositories.CustomerRepo) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) Create(data *models.CustomerModel) (*models.CustomerModel, error) {
	return nil, nil
}

func (s *customerService) GetAll() ([]models.CustomerModel, error) {
	return nil, nil
}

func (s *customerService) GetById(id int) (*models.CustomerModel, error) {
	return nil, nil
}

func (s *customerService) Update(data *models.CustomerModel) error {
	return nil
}

func (s *customerService) Delete(id int) error {
	return nil
}

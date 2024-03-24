package services

import (
	"fmt"
	"strconv"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/repositories"
)

type CustomerService interface {
	Create(data *models.CustomerModel) (*models.CustomerModel, error)
	GetAll() ([]models.CustomerModel, error)
	GetById(string) (*models.CustomerModel, error)
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
	newCust, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}

	return newCust, err
}

func (s *customerService) GetAll() ([]models.CustomerModel, error) {
	customers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return nil, fmt.Errorf("tidak ada data")
	}

	return customers, nil
}

func (s *customerService) GetById(idStr string) (*models.CustomerModel, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	customer, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *customerService) Update(data *models.CustomerModel) error {
	return nil
}

func (s *customerService) Delete(id int) error {
	return nil
}

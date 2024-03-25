package services

import (
	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/repositories"
)

type ProductService interface {
	Create(*models.ProductModel) (*models.ProductModel, error)
	GetAll() ([]models.ProductModel, error)
}

type productService struct {
	repo repositories.ProductRepo
}

func NewProductService(repo repositories.ProductRepo) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(data *models.ProductModel) (*models.ProductModel, error) {
	newProduct, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return newProduct, err
}

func (s *productService) GetAll() ([]models.ProductModel, error) {
	return nil, nil
}

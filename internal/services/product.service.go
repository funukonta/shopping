package services

import (
	"fmt"
	"strings"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/repositories"
)

type ProductService interface {
	Create(*models.ProductModel) (*models.ProductModel, error)
	GetAll() ([]models.ProductModel, error)
	GetByCategory(string) ([]models.ProductModel, error)
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
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("tidak ada data product")
	}

	return products, err
}

func (s *productService) GetByCategory(cat string) ([]models.ProductModel, error) {
	cat = strings.ToUpper(cat)
	products, err := s.repo.GetByCategory(cat)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("tidak ada data product")
	}

	return products, err
}

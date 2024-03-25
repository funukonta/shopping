package services

import (
	"fmt"
	"strconv"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/repositories"
)

type ShopCartService interface {
	GetAllList(id_cust string) (*models.ResJoinCart, error)
	AddCart(string, *models.ShoppingCartDetail) error
	DeleteProduct(string, []models.ShoppingCartDetail) error
	UpdateProduct(*models.ShoppingCart, []models.ShoppingCartDetail) error
}

type shopCartService struct {
	repo repositories.ShopCartRepo
}

func NewShoppingCartService(repo repositories.ShopCartRepo) ShopCartService {
	return &shopCartService{repo: repo}
}

func (s *shopCartService) GetAllList(id_cust string) (*models.ResJoinCart, error) {
	custId, err := strconv.Atoi(id_cust)
	if err != nil {
		return nil, err
	}

	cartId, shopList, err := s.repo.GetAllList(custId)
	if err != nil {
		return nil, err
	}
	if cartId == 0 {
		return nil, fmt.Errorf("tidak ada data")
	}

	shopping_cart := models.ResJoinCart{}

	shopping_cart.Cart.CustID = custId
	shopping_cart.Cart.CartID = cartId
	shopping_cart.Dtl = shopList

	return &shopping_cart, err
}

func (s *shopCartService) AddCart(id_cust string, product *models.ShoppingCartDetail) error {
	custId, err := strconv.Atoi(id_cust)
	if err != nil {
		return err
	}

	if product.Quantity <= 0 {
		return fmt.Errorf("quantity tidak boleh kurang dari 0")
	}

	err = s.repo.AddCart(custId, product)
	return err
}

func (s *shopCartService) DeleteProduct(id_cust string, product []models.ShoppingCartDetail) error {
	custId, err := strconv.Atoi(id_cust)
	if err != nil {
		return err
	}

	err = s.repo.DeleteProduct(custId, product)
	return err
}
func (s *shopCartService) UpdateProduct(*models.ShoppingCart, []models.ShoppingCartDetail) error {
	return nil
}

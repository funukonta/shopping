package handlers

import (
	"net/http"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/services"
	"github.com/funukonta/shopping/pkg"
)

type ShopCartHandler interface {
	GetAllList(w http.ResponseWriter, r *http.Request) error
	AddCart(w http.ResponseWriter, r *http.Request) error
	DeleteProduct(w http.ResponseWriter, r *http.Request) error
	UpdateProduct(w http.ResponseWriter, r *http.Request) error
}

type shopCartHandler struct {
	serv services.ShopCartService
}

func NewShoppingCartHandler(serv services.ShopCartService) ShopCartHandler {
	return &shopCartHandler{serv: serv}
}

func (h *shopCartHandler) GetAllList(w http.ResponseWriter, r *http.Request) error {
	id_cust := r.PathValue("cust")

	shopping_cart, err := h.serv.GetAllList(id_cust)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: shopping_cart, Message: "Berhasil ambil shopping list"}).Send(w)

	return nil
}

func (h *shopCartHandler) AddCart(w http.ResponseWriter, r *http.Request) error {
	id_cust := r.PathValue("cust")
	product := models.ShoppingCartDetail{}
	if err := pkg.GetJsonBody(r, &product); err != nil {
		return err
	}

	err := h.serv.AddCart(id_cust, &product)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: "Berhasil insert to cart"}).Send(w)
	return nil
}
func (h *shopCartHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) error {
	id_cust := r.PathValue("cust")
	product := []models.ShoppingCartDetail{}
	if err := pkg.GetJsonBody(r, &product); err != nil {
		return err
	}

	err := h.serv.DeleteProduct(id_cust, product)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: "Berhasil delete to cart"}).Send(w)
	return nil
}
func (h *shopCartHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) error {
	id_cust := r.PathValue("cust")
	product := []models.ShoppingCartDetail{}
	if err := pkg.GetJsonBody(r, &product); err != nil {
		return err
	}

	err := h.serv.UpdateProduct(id_cust, product)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: "Berhasil update cart"}).Send(w)
	return nil
}

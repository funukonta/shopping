package handlers

import (
	"net/http"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/services"
	"github.com/funukonta/shopping/pkg"
)

type ProductHandler interface {
	Create(w http.ResponseWriter, r *http.Request) error
	GetAll(w http.ResponseWriter, r *http.Request) error
	GetByCategory(w http.ResponseWriter, r *http.Request) error
}

type productHandler struct {
	serv services.ProductService
}

func NewProductHandler(serv services.ProductService) ProductHandler {
	return &productHandler{serv: serv}
}

func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) error {
	product := new(models.ProductModel)
	if err := pkg.GetJsonBody(r, product); err != nil {
		return err
	}

	newProdct, err := h.serv.Create(product)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: newProdct, Message: "Berhasil create product"}).Send(w)
	return nil
}

func (h *productHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	products, err := h.serv.GetAll()
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: products, Message: "Berhasil get products"}).Send(w)
	return nil
}

func (h *productHandler) GetByCategory(w http.ResponseWriter, r *http.Request) error {
	cat := r.PathValue("category")
	products, err := h.serv.GetByCategory(cat)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: products, Message: "Berhasil get products"}).Send(w)
	return nil
}

package handlers

import (
	"net/http"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/services"
	"github.com/funukonta/shopping/pkg"
)

type CustomerHandler interface {
	Create(w http.ResponseWriter, r *http.Request) error
	GetAll(w http.ResponseWriter, r *http.Request) error
	GetById(w http.ResponseWriter, r *http.Request) error
	Update(w http.ResponseWriter, r *http.Request) error
	Delete(w http.ResponseWriter, r *http.Request) error
}

type customerHandler struct {
	serv services.CustomerService
}

func NewCustomerHandler(serv services.CustomerService) CustomerHandler {
	return &customerHandler{serv: serv}
}

func (h *customerHandler) Create(w http.ResponseWriter, r *http.Request) error {
	cust := new(models.CustomerModel)
	if err := pkg.GetJsonBody(r, cust); err != nil {
		return err
	}

	newCust, err := h.serv.Create(cust)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: newCust, Message: "Customer Created!"}).Send(w)
	return nil
}

func (h *customerHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (h *customerHandler) GetById(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (h *customerHandler) Update(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (h *customerHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}

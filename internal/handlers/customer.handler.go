package handlers

import (
	"fmt"
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
	customers, err := h.serv.GetAll()
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: customers, Message: "Berhasil ambil data customers"}).Send(w)
	return nil
}
func (h *customerHandler) GetById(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")
	customer, err := h.serv.GetById(id)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: customer, Message: fmt.Sprintf("Berhasil ambil customer : %s", id)}).Send(w)
	return nil
}
func (h *customerHandler) Update(w http.ResponseWriter, r *http.Request) error {
	cust := new(models.CustomerModel)
	if err := pkg.GetJsonBody(r, cust); err != nil {
		return err
	}

	err := h.serv.Update(cust)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: fmt.Sprintf("Berhasil update customer : %d", cust.ID)}).Send(w)
	return nil
}
func (h *customerHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")
	err := h.serv.Delete(id)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Message: fmt.Sprintf("Berhasil delete customer : %s", id)}).Send(w)
	return nil
}

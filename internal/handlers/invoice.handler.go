package handlers

import (
	"net/http"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/services"
	"github.com/funukonta/shopping/pkg"
)

type InvoiceHandler interface {
	Create(w http.ResponseWriter, r *http.Request) error
	GetInvoice(w http.ResponseWriter, r *http.Request) error
	GetInvoiceById(w http.ResponseWriter, r *http.Request) error
}

type invoiceHandler struct {
	serv services.InvoiceService
}

func NewInvoiceHandler(serv services.InvoiceService) InvoiceHandler {
	return &invoiceHandler{serv: serv}
}

func (h *invoiceHandler) Create(w http.ResponseWriter, r *http.Request) error {
	id_cust := r.PathValue("cust")
	body := []models.InvoiceDetail{}
	if err := pkg.GetJsonBody(r, &body); err != nil {
		return err
	}

	result, err := h.serv.Create(id_cust, body)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: result, Message: "Berhasil create invoice"}).Send(w)
	return nil
}
func (h *invoiceHandler) GetInvoice(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (h *invoiceHandler) GetInvoiceById(w http.ResponseWriter, r *http.Request) error {
	return nil
}

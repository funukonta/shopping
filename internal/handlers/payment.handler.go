package handlers

import (
	"net/http"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/services"
	"github.com/funukonta/shopping/pkg"
)

type PaymentHandler interface {
	DoPayment(w http.ResponseWriter, r *http.Request) error
}

type paymentHandler struct {
	serv services.PaymentService
}

func NewPaymentHandler(serv services.PaymentService) PaymentHandler {
	return &paymentHandler{serv: serv}
}

func (h *paymentHandler) DoPayment(w http.ResponseWriter, r *http.Request) error {
	id_inv := r.PathValue("inv")
	payment := new(models.PaymentModel)
	if err := pkg.GetJsonBody(r, payment); err != nil {
		return err
	}

	result, err := h.serv.DoPayment(id_inv, payment)
	if err != nil {
		return err
	}

	pkg.Response(http.StatusOK, &pkg.JsonBod{Data: result, Message: "SUKSES PAYMENT"}).Send(w)
	return nil
}

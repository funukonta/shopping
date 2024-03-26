package routes

import (
	"net/http"

	"github.com/funukonta/shopping/pkg"
	"github.com/jmoiron/sqlx"
)

func Routes(mux *http.ServeMux, db *sqlx.DB) {
	CustomerRoutes(mux, db)
	ProductRoutes(mux, db)
	ShoopingCartRoutes(mux, db)
	InvoiceRoute(mux, db)
	PaymentRoutes(mux, db)
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		pkg.Response(400, &pkg.JsonBod{Message: err.Error()}).Send(w)
	}
}

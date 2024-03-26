package routes

import (
	"net/http"

	"github.com/funukonta/shopping/internal/handlers"
	"github.com/funukonta/shopping/internal/repositories"
	"github.com/funukonta/shopping/internal/services"
	"github.com/jmoiron/sqlx"
)

func InvoiceRoute(mux *http.ServeMux, db *sqlx.DB) {
	repo := repositories.NewInvoiceRepo(db)
	serv := services.NewInvoiceService(repo)
	invHandler := handlers.NewInvoiceHandler(serv)

	mux.Handle("POST /invoices/{cust}", Handler(invHandler.Create))
	mux.Handle("GET /invoices/{cust}", Handler(invHandler.GetInvoice))
	mux.Handle("GET /invoices/{cust}/{inv}", Handler(invHandler.GetInvoiceById))
}

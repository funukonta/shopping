package routes

import (
	"net/http"

	"github.com/funukonta/shopping/internal/handlers"
	"github.com/funukonta/shopping/internal/repositories"
	"github.com/funukonta/shopping/internal/services"
	"github.com/jmoiron/sqlx"
)

func CustomerRoutes(mux *http.ServeMux, db *sqlx.DB) {
	repo := repositories.NewCustomerRepo(db)
	service := services.NewCustomerService(repo)
	customerHandler := handlers.NewCustomerHandler(service)

	mux.Handle("POST /customers", Handler(customerHandler.Create))
	mux.Handle("GET /customers", Handler(customerHandler.GetAll))
	mux.Handle("PUT /customers/{id}", Handler(customerHandler.Update))
	mux.Handle("DELETE /customers/{id}", Handler(customerHandler.Delete))
	mux.Handle("GET /customers/{id}", Handler(customerHandler.GetById))
}

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

	mux.Handle("POST /register", Handler(customerHandler.Create))

	mux.Handle("POST /customers", authMiddleware(Handler(customerHandler.Create)))
	mux.Handle("GET /customers", authMiddleware(Handler(customerHandler.GetAll)))
	mux.Handle("PUT /customers/{id}", authMiddleware(Handler(customerHandler.Update)))
	mux.Handle("DELETE /customers/{id}", authMiddleware(Handler(customerHandler.Delete)))
	mux.Handle("GET /customers/{id}", authMiddleware(Handler(customerHandler.GetById)))
}

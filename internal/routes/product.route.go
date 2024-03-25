package routes

import (
	"net/http"

	"github.com/funukonta/shopping/internal/handlers"
	"github.com/funukonta/shopping/internal/repositories"
	"github.com/funukonta/shopping/internal/services"
	"github.com/jmoiron/sqlx"
)

func ProductRoutes(mux *http.ServeMux, db *sqlx.DB) {
	repo := repositories.NewProductRepo(db)
	serv := services.NewProductService(repo)
	prodHandler := handlers.NewProductHandler(serv)

	mux.Handle("POST /products", Handler(prodHandler.Create))
	mux.Handle("GET /products", Handler(prodHandler.GetAll))
	mux.Handle("GET /products/{category}", Handler(prodHandler.GetByCategory))

}

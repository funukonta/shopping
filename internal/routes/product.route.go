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

	mux.Handle("POST /products", authMiddleware(Handler(prodHandler.Create)))
	mux.Handle("GET /products", authMiddleware(Handler(prodHandler.GetAll)))
	mux.Handle("GET /products/{category}", authMiddleware(Handler(prodHandler.GetByCategory)))

}

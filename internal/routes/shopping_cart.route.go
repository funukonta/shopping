package routes

import (
	"net/http"

	"github.com/funukonta/shopping/internal/handlers"
	"github.com/funukonta/shopping/internal/repositories"
	"github.com/funukonta/shopping/internal/services"
	"github.com/jmoiron/sqlx"
)

func ShoopingCartRoutes(mux *http.ServeMux, db *sqlx.DB) {
	repo := repositories.NewShoppingCartRepo(db)
	serv := services.NewShoppingCartService(repo)
	shopcartHandler := handlers.NewShoppingCartHandler(serv)

	mux.Handle("GET /shopping-carts/{cust}", Handler(shopcartHandler.GetAllList))
	mux.Handle("POST /shopping-carts/{cust}", Handler(shopcartHandler.AddCart))
	mux.Handle("DELETE /shopping-carts/{cust}", Handler(shopcartHandler.DeleteProduct))
}

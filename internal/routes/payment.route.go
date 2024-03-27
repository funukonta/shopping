package routes

import (
	"net/http"

	"github.com/funukonta/shopping/internal/handlers"
	"github.com/funukonta/shopping/internal/repositories"
	"github.com/funukonta/shopping/internal/services"
	"github.com/jmoiron/sqlx"
)

func PaymentRoutes(mux *http.ServeMux, db *sqlx.DB) {
	repo := repositories.NewPaymentRepo(db)
	serv := services.NewPaymentService(repo)
	paymentHandler := handlers.NewPaymentHandler(serv)

	mux.Handle("POST /payment/{inv}", authMiddleware(Handler(paymentHandler.DoPayment)))
}

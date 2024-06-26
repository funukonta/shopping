package routes

import (
	"net/http"
	"os"

	"github.com/funukonta/shopping/internal/handlers"
	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/repositories"
	"github.com/funukonta/shopping/internal/services"
	"github.com/funukonta/shopping/pkg"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

func Routes(mux *http.ServeMux, db *sqlx.DB) {
	repo := repositories.NewAuthRepo(db)
	serv := services.NewAuthService(repo)
	authHandler := handlers.NewAuthHandler(serv)

	mux.Handle("POST /login", Handler(authHandler.Login))
	mux.Handle("POST /logout", Handler(authHandler.Logout))

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

type Middle func(next http.Handler) http.Handler

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			pkg.Response(http.StatusUnauthorized, &pkg.JsonBod{Message: "bad request"}).Send(w)
			return
		}

		tokenString := cookie.Value
		claim := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				pkg.Response(http.StatusUnauthorized, &pkg.JsonBod{Message: "Unautorized"}).Send(w)
				return
			}
			pkg.Response(http.StatusUnauthorized, &pkg.JsonBod{Message: "bad request"}).Send(w)
			return
		}
		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

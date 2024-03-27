package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/funukonta/shopping/internal/routes"
	"github.com/funukonta/shopping/pkg"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mux := http.NewServeMux()
	time.Sleep(time.Second * 5)
	err := pkg.ConnectAndCreateDB()
	if err != nil {
		log.Println("ERRRRRROOOOOORRR", err)
		return
	}
	db, err := pkg.ConnectPostgre()
	if err != nil {
		log.Panic(err.Error())
		return
	}

	routes.Routes(mux, db)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello")
	})

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	log.Println("API Server listening to port", port)
	http.ListenAndServe(port, mux)
}

// go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
// migrate create -ext sql -dir migrate crate_table_customer
// migrate create -ext sql -dir migrate crate_table_product
// migrate create -ext sql -dir migrate crate_table_cart
// migrate create -ext sql -dir migrate crate_table_invoice
// migrate create -ext sql -dir migrate crate_table_payment

// docker postges
// docker run --name shopping -e POSTGRES_PASSWORD=shopping -p 5432:5432 -d postgres && sleep 2 && docker exec -it shopping psql -U postgres -d postgres -c "CREATE DATABASE shoppingdb;"

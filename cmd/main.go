package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello")
	})

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	http.ListenAndServe(port, mux)
}

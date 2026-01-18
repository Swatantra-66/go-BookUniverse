package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Swatantra-66/go-bookstore/pkg/config"
	"github.com/Swatantra-66/go-bookstore/pkg/middleware"
	"github.com/Swatantra-66/go-bookstore/pkg/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println("--------------------------------------------------")
	fmt.Println("📂 SERVER RUNNING FROM:", dir)
	fmt.Println("--------------------------------------------------")

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, assuming environment variables are set.")
	}

	config.Connect()
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	routes.RegisterBookRoutes(router)

	http.Handle("/", router)
	fmt.Printf("Server running on http://localhost:8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}

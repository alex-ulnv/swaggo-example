package main

import (
	"fmt"
	"log"
	"net/http"

	_ "swaggo-example/docs" // Include generated Swagger docs
	"swaggo-example/api"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Example Accounts API
// @version 1.0
// @description An example of how to document Go API with Swagger (Swaggo).
// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Register routes
	http.HandleFunc("/api/v1/account/", api.GetAccount)
	http.HandleFunc("/api/v1/account", api.CreateAccount)
	http.HandleFunc("/api/v1/update-account", api.UpdateAccount)

	// Swagger endpoint
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

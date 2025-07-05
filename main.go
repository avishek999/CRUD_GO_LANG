package main

import (
	"log"
	"net/http"
	"GO_CRUD/routes"
)

func main() {
	routes.RegisterRoutes()
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

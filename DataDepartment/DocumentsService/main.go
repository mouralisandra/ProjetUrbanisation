package main

import (
	"log"
	"net/http"
	"DocumentsService/documents"
	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Initialize and setup routes for both services
	document.Init()
	document.SetupRoutes(r)

	// Start server
	log.Fatal(http.ListenAndServe(":7000", r))
}

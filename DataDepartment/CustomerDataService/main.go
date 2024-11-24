package main

import (
	"log"
	"net/http"
	"CustomerDataService/customerdata"
	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Initialize and setup routes for both services
	data.Init()
	data.SetupRoutes(r)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

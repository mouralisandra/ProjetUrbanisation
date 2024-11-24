package document

import (
	"github.com/gorilla/mux"
)


// Function to initialize the Document service (optional)
func Init() {
	// Pre-populating with some document
	documents = append(documents, Document{ID: "1", Isbn: "438227", Title: "Content of Document1", Author: &Author{Firstname: "Sandra", Lastname: "Mourali"}})
	documents = append(documents, Document{ID: "2", Isbn: "454555", Title: "Content of Document2", Author: &Author{Firstname: "Nada", Lastname: "Mankai"}})
}

// Function to setup routes for the Document service
func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/documents", GetDocuments).Methods("GET")
	r.HandleFunc("/documents/{id}", GetDocument).Methods("GET")
	r.HandleFunc("/documents", CreateDocument).Methods("POST")
	r.HandleFunc("/documents/{id}", UpdateDocument).Methods("PUT")
	r.HandleFunc("/documents/{id}", DeleteDocument).Methods("DELETE")
}

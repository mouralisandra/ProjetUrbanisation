package document

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// In-memory document store
var documents []Document

// Get all documents
func GetDocuments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(documents)
}

// Get single document
func GetDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range documents {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Document{})
}

// Create new document
func CreateDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var document Document
	_ = json.NewDecoder(r.Body).Decode(&document)
	document.ID = strconv.Itoa(rand.Intn(100000000))
	documents = append(documents, document)
	json.NewEncoder(w).Encode(document)
}

// Update document
func UpdateDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range documents {
		if item.ID == params["id"] {
			documents = append(documents[:index], documents[index+1:]...)
			var document Document
			_ = json.NewDecoder(r.Body).Decode(&document)
			document.ID = params["id"]
			documents = append(documents, document)
			json.NewEncoder(w).Encode(document)
			return
		}
	}
}

// Delete document
func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range documents {
		if item.ID == params["id"] {
			documents = append(documents[:index], documents[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(documents)
}

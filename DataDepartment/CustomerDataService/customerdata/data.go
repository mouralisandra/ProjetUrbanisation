package data

import (
	"github.com/gorilla/mux"
)


// Function to initialize the Data service (optional)
func Init() {
	// Pre-populating with some data
	datas = append(datas, Data{ID: "1", Isbn: "438227", Title: "Data1", Author: &Author{Firstname: "Sandra", Lastname: "Mourali"}})
	datas = append(datas, Data{ID: "2", Isbn: "454555", Title: "Data2", Author: &Author{Firstname: "Nada", Lastname: "Mankai"}})
}

// Function to setup routes for the Data service
func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/datas", GetDatas).Methods("GET")
	r.HandleFunc("/datas/{id}", GetData).Methods("GET")
	r.HandleFunc("/datas", CreateData).Methods("POST")
	r.HandleFunc("/datas/{id}", UpdateData).Methods("PUT")
	r.HandleFunc("/datas/{id}", DeleteData).Methods("DELETE")
}

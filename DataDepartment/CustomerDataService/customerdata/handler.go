package data

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// In-memory data store
var datas []Data

// Get all datas
func GetDatas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datas)
}

// Get single data
func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, item := range datas {
		if item.ID == id {
			// If ID is found, encode the item and return
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// If ID is not found, return 400
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
}


// Create new data
func CreateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	data.ID = strconv.Itoa(rand.Intn(100000000))
	datas = append(datas, data)
	json.NewEncoder(w).Encode(data)
}

// Update data
func UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range datas {
		if item.ID == params["id"] {
			datas = append(datas[:index], datas[index+1:]...)
			var data Data
			_ = json.NewDecoder(r.Body).Decode(&data)
			data.ID = params["id"]
			datas = append(datas, data)
			json.NewEncoder(w).Encode(data)
			return
		}
	}
}

// Delete data
func DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range datas {
		if item.ID == params["id"] {
			datas = append(datas[:index], datas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(datas)
}

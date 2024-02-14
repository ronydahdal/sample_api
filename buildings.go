package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// define buildings struct here
type BuildingsHandler struct {
}

// define CRUD operations
// take Response and Request as parameter to respond and read requests

// List buildings will just use list buildings function to return buildings lsit
func (b BuildingsHandler) ListBuildings(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(listBuildings())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

// reads requested building id from user, which then gets passed to get building
func (b BuildingsHandler) GetBuildings(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	building := getBuilding(id)
	// check for some errors before returning json
	if building == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(building)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BuildingsHandler) CreateBuildings(w http.ResponseWriter, r *http.Request) {
	var building Building
	// decode the requested new building
	err := json.NewDecoder(r.Body).Decode(&building)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// store the building if it's not a bad request
	storeBuilding(building)

	// encode new building
	err = json.NewEncoder(w).Encode(building)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BuildingsHandler) UpdateBuildings(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var building Building
	err := json.NewDecoder(r.Body).Decode(&building)
	// catch if request isn't bad
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// update building with id
	updatedBuilding := updateBuilding(id, building)
	if updatedBuilding == nil {
		http.Error(w, "Building not found", http.StatusNotFound)
		return
	}
	// encode json
	err = json.NewEncoder(w).Encode(updatedBuilding)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BuildingsHandler) DeleteBuildings(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// delete the building with given id
	building := deleteBuilding(id)
	// check if nil was returned
	if building == nil {
		http.Error(w, "Building not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

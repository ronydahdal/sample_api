package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// define CRUD operations
// take Response and Request as parameter to respond and read requests
// func (b BuildingsHandler) ListBuildings(w http.ResponseWriter, r *http.Request)   {}
// func (b BuildingsHandler) GetBuildings(w http.ResponseWriter, r *http.Request)    {}
// func (b BuildingsHandler) CreateBuildings(w http.ResponseWriter, r *http.Request) {}
// func (b BuildingsHandler) UpdateBuildings(w http.ResponseWriter, r *http.Request) {}
// func (b BuildingsHandler) DeleteBuildings(w http.ResponseWriter, r *http.Request) {}

func main() {
	// new router instance
	r := chi.NewRouter()
	// logging middlewarre
	r.Use(middleware.Logger)
	// set up route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// return OK to client
		w.Write([]byte("OK"))
	})
	// mount route to main route -- this is to link to root path which is /buildings
	r.Mount("/buildings", BuildingsRoutes())
	// start a server to listen at port 3000
	http.ListenAndServe(":3000", r)
}

func BuildingsRoutes() chi.Router {
	// initialize new router to use for API operations
	r := chi.NewRouter()
	// attach routes to BuildingsHandler
	buildingsHandler := BuildingsHandler{}
	// use chi functions for CRUD
	r.Get("/", buildingsHandler.ListBuildings)
	r.Post("/", buildingsHandler.CreateBuildings)
	r.Get("/{id}", buildingsHandler.GetBuildings)
	r.Put("/{id}", buildingsHandler.UpdateBuildings)
	r.Delete("/{id}", buildingsHandler.DeleteBuildings)
	return r
}

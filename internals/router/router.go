package router

import (
	"exoplanet-service/internals/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", handlers.AddExoplanet).Methods("POST")
	r.HandleFunc("/exoplanets", handlers.ListExoplanets).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.GetExoplanetByID).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.UpdateExoplanet).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", handlers.DeleteExoplanet).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel", handlers.FuelEstimation).Methods("GET")
	return r
}

package handlers

import (
	"encoding/json"
	"exoplanet-service/internals"
	validation "exoplanet-service/pkg/vlaidation"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var exoplanets = map[string]internals.Exoplanet{}

func AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet internals.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validation.ValidateExoplanet(exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exoplanet.ID = strconv.Itoa(len(exoplanets) + 1)
	exoplanets[exoplanet.ID] = exoplanet
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}

func ListExoplanets(w http.ResponseWriter, r *http.Request) {
	planets := []internals.Exoplanet{}
	for _, exoplanet := range exoplanets {
		planets = append(planets, exoplanet)
	}
	json.NewEncoder(w).Encode(planets)
}

func GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	exoplanet, exists := exoplanets[id]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	_, exists := exoplanets[id]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	var updatedExoplanet internals.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&updatedExoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validation.ValidateExoplanet(updatedExoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedExoplanet.ID = id
	exoplanets[id] = updatedExoplanet
	json.NewEncoder(w).Encode(updatedExoplanet)
}

func DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if _, exists := exoplanets[id]; !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	delete(exoplanets, id)
	w.WriteHeader(http.StatusNoContent)
}

func FuelEstimation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	crewCapacityStr := r.URL.Query().Get("crew_capacity")
	if crewCapacityStr == "" {
		http.Error(w, "Crew capacity is required", http.StatusBadRequest)
		return
	}
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}
	exoplanet, exists := exoplanets[id]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	var gravity float64
	if exoplanet.Type == internals.GasGiant {
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	} else if exoplanet.Type == internals.Terrestrial {
		gravity = *exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	}
	fuel := float64(exoplanet.Distance) / (gravity * gravity) * float64(crewCapacity)
	json.NewEncoder(w).Encode(map[string]float64{"fuel_estimation": fuel})
}

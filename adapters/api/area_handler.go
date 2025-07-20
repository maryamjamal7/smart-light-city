package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/service"
)

func RegisterAreaRoutes(r *mux.Router, areaSvc *service.AreaService) {
	r.HandleFunc("/areas", createAreaHandler(areaSvc)).Methods("POST")
	r.HandleFunc("/cities", listCitiesHandler(areaSvc)).Methods("GET")
	r.HandleFunc("/cities/{id}/zones", listZonesHandler(areaSvc)).Methods("GET")
}

func createAreaHandler(areaSvc *service.AreaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var area model.Area
		if err := json.NewDecoder(r.Body).Decode(&area); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := areaSvc.CreateArea(r.Context(), &area); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(area)
	}
}

func listCitiesHandler(areaSvc *service.AreaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cities, err := areaSvc.ListCities(r.Context())
		if err != nil {
			http.Error(w, "Error fetching cities", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(cities)
	}
}

func listZonesHandler(areaSvc *service.AreaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cityIDStr := mux.Vars(r)["id"]
		cityID, err := strconv.Atoi(cityIDStr)
		if err != nil {
			http.Error(w, "Invalid city ID", http.StatusBadRequest)
			return
		}

		zones, err := areaSvc.ListZonesByCityID(r.Context(), uint(cityID))
		if err != nil {
			http.Error(w, "Error fetching zones", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(zones)
	}
}

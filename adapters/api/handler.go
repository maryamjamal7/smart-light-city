package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/maryamjamal7/smart-light-city/domain/service"
)

// RegisterAllRoutes wires all routes using Gorilla Mux
func RegisterAllRoutes(areaSvc *service.AreaService, lumiereSvc *service.LumiereService, cmdSvc *service.CommandService, manager *service.CityManager) *mux.Router {
	r := mux.NewRouter()

	// Route groups
	RegisterAreaRoutes(r, areaSvc)
	RegisterLumiereRoutes(r, lumiereSvc)
	RegisterCommandRoutes(r, cmdSvc)

	// City-wide control
	r.HandleFunc("/city/poweroff", HandleCityPowerOff(manager)).Methods("POST")
	r.HandleFunc("/city/dim", HandleCityScheduleDim(manager)).Methods("POST")

	return r
}

func HandleCityPowerOff(manager *service.CityManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := manager.PowerOffAll(r.Context())
		if err != nil {
			http.Error(w, "Failed to power off city: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All lights powered off"))
	}
}

type DimRequest struct {
	Dim int       `json:"dim"`
	At  time.Time `json:"at"`
}

func HandleCityScheduleDim(manager *service.CityManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DimRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err := manager.ScheduleDimForAll(r.Context(), req.Dim, req.At)
		if err != nil {
			http.Error(w, "Failed to schedule dimming: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Dimming scheduled"))
	}
}

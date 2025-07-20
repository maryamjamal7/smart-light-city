package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/service"
)

func RegisterLumiereRoutes(r *mux.Router, svc *service.LumiereService) {
	r.HandleFunc("/lumiere", createLumiereHandler(svc)).Methods("POST")
	r.HandleFunc("/lumiere/{id}", updateLumiereStateHandler(svc)).Methods("PUT")
	r.HandleFunc("/area/{area_id}/lumiere", listByAreaHandler(svc)).Methods("GET")
}

func createLumiereHandler(svc *service.LumiereService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var l model.Lumiere
		if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if err := svc.Create(r.Context(), &l); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(l)
	}
}

func updateLumiereStateHandler(svc *service.LumiereService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var payload struct {
			Power bool `json:"power"`
			Dim   int  `json:"dim"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if err := svc.UpdateState(r.Context(), uint(id), payload.Power, payload.Dim); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func listByAreaHandler(svc *service.LumiereService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		areaIDStr := mux.Vars(r)["area_id"]
		areaID, err := strconv.Atoi(areaIDStr)
		if err != nil {
			http.Error(w, "Invalid area ID", http.StatusBadRequest)
			return
		}

		list, err := svc.ListByArea(r.Context(), uint(areaID))
		if err != nil {
			http.Error(w, "Failed to fetch", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(list)
	}
}

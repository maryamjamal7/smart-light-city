package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/service"
)

// RegisterCommandRoutes wires up the command-related routes
func RegisterCommandRoutes(r *mux.Router, svc *service.CommandService) {
	r.HandleFunc("/commands", createCommandHandler(svc)).Methods("POST")
	r.HandleFunc("/commands", listCommandsHandler(svc)).Methods("GET")
}

// Handles POST /commands to schedule a command
func createCommandHandler(svc *service.CommandService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd model.Command

		// Decode JSON body into Command model
		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Validate and save via service
		if err := svc.ScheduleCommand(r.Context(), &cmd); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond with the saved command
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(cmd)
	}
}

// Handles GET /commands to list all scheduled commands
func listCommandsHandler(svc *service.CommandService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commands, err := svc.ListCommands(r.Context())
		if err != nil {
			http.Error(w, "Failed to list commands", http.StatusInternalServerError)
			return
		}

		// Return all commands as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(commands)
	}
}

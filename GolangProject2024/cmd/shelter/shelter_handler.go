// shelter_handler.go
package main

import (
	"encoding/json"
	"net/http"
	"strconv" // Import strconv package for string conversion

	"github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	response := map[string]string{"error": message}
	respondWithJSON(w, code, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) getShelterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shelterID := vars["shelterId"]

	// Convert shelterID to an integer using the convertToInt function
	id, err := convertToInt(shelterID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid shelter ID")
		return
	}

	// Pass the integer ID to app.models.Shelters.Get
	shelter, err := app.models.Shelters.Get(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Shelter not found")
		return
	}

	respondWithJSON(w, http.StatusOK, shelter)
}

// Function to convert string to integer
func convertToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

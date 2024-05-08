package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/po133na/go-mid/pkg/shelter/model"
)

func (app *application) createVolunteerHandler(w http.ResponseWriter, r *http.Request) {
	var input model.Volunteer

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = app.models.Volunteer.Insert(&input)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, input)
}

func (app *application) getVolunteerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volunteerID := vars["volunteerId"]

	volunteer, err := app.models.Volunteers.GetVolunteer(volunteerID)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "Volunteer not found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, volunteer)
}

func (app *application) updateVolunteerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volunteerID := vars["volunteerId"]

	var input model.Volunteer

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = app.models.Volunteers.UpdateVolunteer(volunteerID, &input)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) deleteVolunteerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volunteerID := vars["volunteerId"]

	err := app.models.Volunteers.DeleteVolunteer(volunteerID)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

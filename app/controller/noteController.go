package controller

import (
	"encoding/json"
	"net/http"

	"github.com/awalludinfajar/note-go-api.git/app/model"
	"github.com/awalludinfajar/note-go-api.git/app/service"
	"github.com/gorilla/mux"
)

func GetNote(w http.ResponseWriter, r *http.Request) {
	notes, err := service.GetAllNote()

	if err != nil {
		http.Error(w, "Failed to load checklists", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var newNote model.Note
	err := json.NewDecoder(r.Body).Decode(&newNote)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdNote, err := service.CreateNote(newNote.Title, newNote.Content)
	if err != nil {
		http.Error(w, "Failed to create note: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdNote)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedNote model.Note
	err := json.NewDecoder(r.Body).Decode(&updatedNote)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	note, err := service.UpdateNote(id, updatedNote.Title, updatedNote.Content)
	if err != nil {
		http.Error(w, "Failed to update note: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := service.DeleteNote(id)
	if err != nil {
		http.Error(w, "Failed to delete note: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

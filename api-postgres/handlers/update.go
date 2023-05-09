package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fhlpassis/golang-postgres/models"
	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error converting id to int: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder((r.Body)).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)

	if rows > 1 {
		log.Printf("Error updating todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any

	if err != nil {
		resp = map[string]any{"error": err.Error()}
	} else {
		resp = map[string]any{"id": id}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

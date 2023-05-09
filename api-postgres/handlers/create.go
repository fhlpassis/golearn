package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fhlpassis/golang-postgres/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder((r.Body)).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{"error": err.Error()}
	} else {
		resp = map[string]any{"id": id}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

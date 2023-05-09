package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fhlpassis/golang-postgres/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error listing todos: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

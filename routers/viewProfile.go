package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Paskual86/go-react-mongodb.git/db"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "The Id is required", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(id)

	if err != nil {
		http.Error(w, "There was an error when we try to search the profile"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

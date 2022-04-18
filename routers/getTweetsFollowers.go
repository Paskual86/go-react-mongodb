package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Paskual86/go-react-mongodb.git/db"
)

func GetTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page parameter is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Page parameter should be numeric", http.StatusBadRequest)
		return
	}

	response, correct := db.GetTweetsFollowers(AuthorizatedUserId, page)

	if !correct {
		http.Error(w, "Error trying to read tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

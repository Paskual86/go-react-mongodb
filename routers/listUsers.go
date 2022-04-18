package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Paskual86/go-react-mongodb.git/db"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUserParameter := r.URL.Query().Get("type")
	pageParameter := r.URL.Query().Get("page")
	searchParameter := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(pageParameter)

	if err != nil {
		http.Error(w, "The parameter Page is required. Shold be an integer greater than 0 (zero)", http.StatusBadRequest)
		return
	}

	page := int64(pageTemp)

	results, status, _ := db.ReadAllUsers(AuthorizatedUserId, page, searchParameter, typeUserParameter)
	if !status {
		http.Error(w, "Error when we try to read the users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(results)
}

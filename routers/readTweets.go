package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Paskual86/go-react-mongodb.git/db"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	userId := r.URL.Query().Get("id")

	if len(userId) < 1 {
		http.Error(w, "The Id is required", http.StatusBadRequest)
		return
	}

	var pageNumber int64 = 1

	pageParameter := r.URL.Query().Get("page")
	if len(pageParameter) > 1 {
		page, errorConvertParameter := strconv.Atoi(pageParameter)
		if errorConvertParameter != nil {
			http.Error(w, "The Page number should be greater than zero", http.StatusBadRequest)
			return
		}
		pageNumber = int64(page)
	}

	response, correct := db.ReadTweets(userId, pageNumber)

	if !correct {
		http.Error(w, "There was an error reading Tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

package routers

import (
	"net/http"

	"github.com/Paskual86/go-react-mongodb.git/db"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	tweetId := r.URL.Query().Get("tweetId")

	if len(tweetId) < 1 {
		http.Error(w, "The Tweet Id is required", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(tweetId, AuthorizatedUserId)

	if err != nil {
		http.Error(w, "There was an error when trying to delete the tweet", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

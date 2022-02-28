package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/models"
)

/*InsertTweet*/
func InsertTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet

	err := json.NewDecoder(r.Body).Decode(&tweet)

	if err != nil {
		http.Error(w, "Error in the input "+err.Error(), 400)
		return
	}

	tweet.UserId = AuthorizatedUserId
	tweet.Date = time.Now()

	if len(tweet.Message) == 0 {
		http.Error(w, "The message is empty", 400)
		return
	}

	_, status, err := db.InsertTweet(tweet)

	if err != nil {
		http.Error(w, "There was an error when we try to insert the record"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "It was not possible insert the record", 400)
		return
	}

	// Aca utilizamos unas constantes declaradas en GO, para indicar que el registro fue creado correctamente.
	w.WriteHeader(http.StatusCreated)
}

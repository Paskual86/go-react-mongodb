package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Invalid input data", 400)
		return
	}

	var status bool

	status, err = db.UpdateUser(t, AuthorizatedUserId)

	if err != nil {
		http.Error(w, "There was some error when we try to update the record. Please retry"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "We could not update the user profile", 400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

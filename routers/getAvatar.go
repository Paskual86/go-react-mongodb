package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"github.com/Paskual86/go-react-mongodb.git/db"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {
	// Get the user Id from the query string
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Missing Parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	file, err := os.Open(constants.AVATAR_FOLDER + profile.Avatar)

	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, file)

	if err != nil {
		http.Error(w, "Error copying image", http.StatusBadRequest)
		return
	}
}

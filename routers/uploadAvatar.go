package routers

import (
	"net/http"

	"github.com/Paskual86/go-react-mongodb.git/constants"
	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/files"
	"github.com/Paskual86/go-react-mongodb.git/models"
	"github.com/Paskual86/go-react-mongodb.git/path"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	// Get the extension
	extension := path.GetExtension(handler.Filename)

	targetFilename := constants.AVATAR_FOLDER + AuthorizatedUserId + "." + extension

	status, err := files.CopyFile(targetFilename, file)

	if err != nil || !status {
		http.Error(w, "Error uploading the file. Details: "+err.Error(), http.StatusBadRequest)
		return
	}

	//Here began the insert in the Database.
	var user models.User
	user.Avatar = AuthorizatedUserId + "." + extension
	status, err = db.UpdateUser(user, AuthorizatedUserId)

	if err != nil || !status {
		http.Error(w, "Error savings information in the database. Details: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

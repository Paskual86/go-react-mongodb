package routers

import (
	"net/http"

	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/models"
	"github.com/Paskual86/go-react-mongodb.git/routers/errors"
)

func InsertRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		errors.MissingParameter(w)
		return
	}

	var t models.Relation
	t.UserId = AuthorizatedUserId
	t.UserRelationId = ID

	status, err := db.InsertRelation(t)

	if err != nil {
		errors.DatabaseError(w, err)
		return
	}

	if !status {
		errors.InsertFailed(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

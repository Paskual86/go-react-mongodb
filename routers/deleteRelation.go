package routers

import (
	"net/http"

	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/models"
	"github.com/Paskual86/go-react-mongodb.git/routers/errors"
)

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "The user Id is required", http.StatusBadRequest)
		return
	}

	var relation models.Relation

	relation.UserRelationId = ID
	relation.UserId = AuthorizatedUserId

	status, err := db.DeleteRelation(relation)

	if err != nil {
		errors.DatabaseErrorDelete(w, err)
		return
	}

	if status == false {
		errors.DeleteFailed(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

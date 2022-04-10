package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/models"
)

func QueryRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var relation models.Relation
	relation.UserId = AuthorizatedUserId
	relation.UserRelationId = ID

	var response models.ResponseQueryRelation

	fmt.Println(relation)
	status, err := db.QueryRelation(relation)
	response.Status = (err == nil && status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

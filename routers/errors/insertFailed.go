package errors

import "net/http"

func InsertFailed(w http.ResponseWriter) {
	http.Error(w, "There was not possible to insert the record", http.StatusBadRequest)
}

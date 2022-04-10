package errors

import "net/http"

func DeleteFailed(w http.ResponseWriter) {
	http.Error(w, "There was not possible to delete the record", http.StatusBadRequest)
}

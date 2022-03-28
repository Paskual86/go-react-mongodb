package errors

import "net/http"

func MissingParameter(w http.ResponseWriter) {
	http.Error(w, "Missing Parameter", http.StatusBadRequest)
	return
}

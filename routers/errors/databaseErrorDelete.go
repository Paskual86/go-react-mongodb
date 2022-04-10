package errors

import "net/http"

func DatabaseErrorDelete(w http.ResponseWriter, detail error) {
	http.Error(w, "There was an error trying to delete the record in the Database. Details"+detail.Error(), http.StatusBadRequest)
}

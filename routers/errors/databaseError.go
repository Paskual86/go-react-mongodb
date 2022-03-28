package errors

import "net/http"

func DatabaseError(w http.ResponseWriter, detail error) {
	http.Error(w, "There was an error trying to insert the record in the Database. Details"+detail.Error(), http.StatusBadRequest)
}

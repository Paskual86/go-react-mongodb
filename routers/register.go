package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	// El Body es un objeto de tipo stream, una vez que se lee se destruye
	// y no se puede volver a leer.
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		// Los codigos de error 400, tienen que ver con la convension.
		// Si los codigos de error estan entre 200 y 400 son excepciones.
		// Los codigos superiores a 400 son errores
		http.Error(w, "Error in the input "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "The password lenght should be have more than 6 chars", 400)
		return
	}

	_, found, _ := db.UsersExists(t.Email)
	if found {
		http.Error(w, "There is another user with the same email", 400)
		return
	}

	_, status, err := db.InsertRecord(t)

	if err != nil {
		http.Error(w, "There was an error when we try to insert the record"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "It was not possible insert the record", 400)
		return
	}

	// Aca utilizamos unas constantes declaradas en GO, para indicar que el registro fue creado correctamente.
	w.WriteHeader(http.StatusCreated)
}

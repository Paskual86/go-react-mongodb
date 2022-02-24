package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Paskual86/go-react-mongodb.git/db"
	"github.com/Paskual86/go-react-mongodb.git/jwt"
	"github.com/Paskual86/go-react-mongodb.git/models"
)

/*Login*/
func Login(w http.ResponseWriter, r *http.Request) {
	// Debemos indicar el formato de salida
	w.Header().Add("content-type", "application/json")

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "User or Password are invalid"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The email is required"+err.Error(), 400)
		return
	}

	usr, exists := db.TryLogin(t.Email, t.Password)

	if !exists {
		http.Error(w, "User or Password are invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(usr)

	if err != nil {
		http.Error(w, "There was an error generating the JWT"+err.Error(), 400)
	}

	result := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

	// Adicional: Como crear una cookie desde el backend

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

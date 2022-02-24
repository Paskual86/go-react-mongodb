package db

import (
	"github.com/Paskual86/go-react-mongodb.git/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {

	usu, found, _ := UsersExists(email)

	if !found {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}

	return usu, true
}

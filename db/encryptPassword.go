package db

import "golang.org/x/crypto/bcrypt"

/*Encrypt Password: It is a function that allow encrypt the password*/
func EncryptPassword(value string) (string, error) {
	// El costo esta relacionado con la funcion de encriptacion. Se interpreta como la cantidad de pasadas que va a tener el
	// texto que se va a encriptar.
	// En este caso el texto se va a encriptar 256 (2^8) veces.
	// Mientras mas grande es el valor, mas va a demorar la encriptacion.
	cost := 8

	// La funcion "GenerateFromPassword" recibe un array de bytes como parametros.
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), cost)
	return string(bytes), err
}

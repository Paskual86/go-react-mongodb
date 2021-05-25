package main

import (
	"fmt"
	"time"

	us "./user"
)

//Como declarar HERENCIA
type pepe struct {
	us.Usuario
}

func main() {
	user := new(pepe)
	user.AltaUsuario(1, "Noah", time.Now(), true)
	fmt.Println(user)
}

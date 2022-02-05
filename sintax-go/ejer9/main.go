package main

import (
	"fmt"
	"time"

	us "example.com/user"
)

//Como declarar HERENCIA
type pepe struct {
	us.Usuario
	Title string
}

func main() {
	user := new(pepe)
	user.AltaUsuario(1, "Noah", time.Now(), true)
	user.Title = "MR"
	fmt.Println(user)
}

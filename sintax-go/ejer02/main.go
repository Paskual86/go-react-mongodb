package main

/*
	Declaracion de variables
*/

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool

func main() {
	numero2, texto2, status2 := 10, "Es texto", false
	var moneda int64 = 50
	numero2 = int(moneda)

	texto = strconv.Itoa(int(moneda))

	fmt.Println("Ejer 02. Variables")

	fmt.Print("numero = ")
	fmt.Println(numero2)
	fmt.Println("texto2 = " + texto2)
	fmt.Println("texto = " + texto)
	fmt.Print("Status 2 = ")
	fmt.Println(status2)
}

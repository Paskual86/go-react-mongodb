package main

import "fmt"

func main() {
	exponencial(2)
}

func exponencial(numero int) {
	if numero > 100000000 {
		return
	}
	fmt.Println(numero)
	exponencial(numero * 2)
}

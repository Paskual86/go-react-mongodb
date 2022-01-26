package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Mostrar y aceptar datos")

	var numero1 int
	var numero2 int

	fmt.Println("Ingrese un numero 1")
	//fmt.Scanf("%d", &numero1)
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese un numero 2")
	//fmt.Scanf("%d", &numero2)
	fmt.Scanln(&numero2)

	scanner := bufio.NewScanner(os.Stdin)

	var leyenda string

	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1+numero2)
}

package main

import "fmt"

func main() {

	fmt.Println("Test Condicionales")

	estado := true

	if estado {
		fmt.Println(estado)
	} else {
		fmt.Println("Es Falso")
	}

	fmt.Println("Ejemplo de Switch")

	switch numero := 10; numero {
	case 1:
		{
			fmt.Println("Valor 1")
		}
	case 2:
		fmt.Println("Valor 2")
	case 3:
		fmt.Println("Valor 3")
	case 4:
		fmt.Println("Valor 4")
	case 5:
		fmt.Println("Valor 5")
	default:
		fmt.Println("Valor Supera el 5")
	}
}

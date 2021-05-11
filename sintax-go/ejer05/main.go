package main

import "fmt"

func main() {

	/*for i := 0; i < count; i++ {

	}*/
	/*
		numero := 0
		for {
			fmt.Println("Continuo")
			fmt.Println("Ingrese el numero secreto")
			fmt.Scanln(&numero)
			if numero == 100 {
				break
			}
		}*/

	/*var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)

		if i == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			i *= 2
			continue // Manda el control al inicio del for.
		}
		fmt.Printf("     Paso por aqui \n")

		i++
	}*/

	var i int = 0
RUTINA:
	for i < 10 {
		if i == 4 {
			i += 2
			fmt.Println("voy a RUTINA sumando 2 a 1")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
}

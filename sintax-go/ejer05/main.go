package main

import "fmt"

func main() {

	var i int = 0

	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)
		i++
	}
	fmt.Println()
	fmt.Println("Forma tradicional del FOR")

	for j := 1; j < 10; j++ {
		fmt.Printf("\n Valor de i: %d", j)
	}

	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)

		if i == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			i *= 2
			continue // Manda el control al inicio del for.
		}
		fmt.Printf("     Paso por aqui \n")

		i++
	}

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

	/**/

	i = 0
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

package main

import "fmt"

func main() {
	/*fmt.Println(uno(5))
	fmt.Println(uno_bis(10))

	numero, estado := dos(12)
	fmt.Println(numero)
	fmt.Println(estado)*/

	fmt.Println(Calculo(1, 2))
	fmt.Println(Calculo(1, 5, 3))
	fmt.Println(Calculo(1, 5, 3, 8))

	fmt.Printf("Sumo 5 + 7 = %d \n", CalculoAnon(5, 7))

	CalculoAnon = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Resto 10 - 4 = %d \n", CalculoAnon(10, 4))

	// CLOSURES, son funciones anonimas.
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	// Cuando se este ejecutando Tabla, se esta ejecutando la funcion. No el Resto.
	// Es como crear una clase e instanciarla. Las variables se setean como en un constructor.
	// Y utilzamos una funcion dentro de la clase, que incremente el valor inicializado en el constructor.
	//
	for i := 1; i < 11; i++ {
		fmt.Printf("%d por 2 = %d\n", i, MiTabla())
	}

}

// Estructura de una funcion
func uno(numero int) int {
	return numero * 2
}

// Le puedo agregar un nombre de salida a la variable
func uno_bis(numero int) (z int) {
	z = numero * 2
	return
}

//Parametros de salida
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// Funcion con parametros de entrada variables
func Calculo(numero ...int) int {
	total := 0
	// Range es una inscruccion que te devuelve dos valores.
	// El primer elemento que devuelve es la ubicacion del elemento en el array.
	// El segundo es el valor.
	/*for _, num := range numero {
		total += num
	}*/
	for item, num := range numero {
		total += num
		fmt.Printf("\n item %d \n", item)
	}
	return total
}

// Funciones Anonimas, una forma de sobreescribir el comportamiento de una funcion.
var CalculoAnon func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}

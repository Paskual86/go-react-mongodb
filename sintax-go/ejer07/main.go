/*
	Exercise about Vectors and Matrix
*/

package main

import "fmt"

// Por default se inicializa en todos 0
var tabla [10]int

var matrix [5][7]int

func main() {
	fmt.Println("Matrix and Vectors")

	fmt.Println("Vector")
	tabla[0] = 5
	tabla[6] = 10
	fmt.Println(tabla)
	// Otra forma de crear e inicializar un vector
	tabla2 := [10]int{5, 6, 9, 8, 7, 8}

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	fmt.Println("Matrix")
	// fmt.Println(matrix)
	// Fila X Columna
	matrix[1][3] = 5
	matrix[2][3] = 8
	matrix[3][3] = 9

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	/*
			Otra forma de recorrer una matriz. Recorre por fila

		for i := range matrix {
			fmt.Println(matrix[i])
		}
	*/

	// SLICE
	// Vectores dinamicos
	// No hay listas en Golang
	// Como se declara
	//matrix2 = []int{2, 7, 8}
	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[2:4]
	//porcion := elementos[:4]
	//porcion := elementos[2:]
	fmt.Println(elementos)
	fmt.Println(porcion)
}

// Con la funcion cap devolvemos la capacidad del array, para este caso el valor devuelto es 20
// con la funcion len devolvemos la cantidad de elementos en el array
func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

// En este ejemplo vamos a ver como crear un vector de forma dinamica.
// La funcion append, permite ingresar elementos al vector
// Lo llamativo de este ejemplo es que a pesar de que la capacidad definida originalmente es 0, go permite incrementar la capacidad en la potencia de 2
// Para este ejemplo la capacidad final va a ser 128 (2^7)
func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}

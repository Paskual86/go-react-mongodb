package main

/*Mapas en GO*/

import "fmt"

func main() {
	paises := make(map[string]string)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	//fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":       39,
		"Real Madrid":     38,
		"Chivas":          37,
		"Almirante Brown": 36}

	// Agrego un nuevo item
	campeonato["River Plate"] = 25
	// Editar un valor del mapa
	campeonato["Chivas"] = 55

	//fmt.Println(campeonato)
	// Como eliminar un elemento
	//delete(campeonato, "Real Madrid")
	//fmt.Println(campeonato)

	/*for team, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d\n", team, puntaje)
	}*/

	puntaje, existe := campeonato["Chicago"]
	fmt.Printf("El puntaje capturado es %d, y el equipo %t \n", puntaje, existe)
}

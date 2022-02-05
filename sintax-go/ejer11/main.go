package main

import "fmt"

/* ----------------------------------- Interface declaration Start ----------------------------------------*/
type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegeta() string
}

/* ----------------------------------- Interface declaration End ------------------------------------------*/

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	estavivo   bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }

func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool { return h.estavivo }

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando\n", hu.sexo())
}

/*********************************Type*****************************************/
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	estavivo   bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.estavivo }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() {
		return 1
	} else {
		return 0
	}
}

/*Ser vivo*/
func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main() {
	/*Pedro := new(hombre)
	Pedro.esHombre = true
	Maria := new(mujer)
	HumanosRespirando(Pedro)
	HumanosRespirando(Maria)*/

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	AnimalesRespirar(Dogo)

	totalCarnivoros += AnimalesCarnivoros(Dogo)
	fmt.Printf("Total carnivoros %d\n", totalCarnivoros)

	fmt.Printf("Estoy vivo = %t", estoyVivo(Dogo))
}

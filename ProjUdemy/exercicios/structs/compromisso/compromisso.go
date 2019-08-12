package main

import "fmt"

type horario struct {
	horaEspecifica string
}

type data struct {
	dataEspecifica string
}

type compromisso struct {
	data
	horario
	// Ambos compostos por structs separadas
	//podendo ser ultilizadas dentro dessa mesma
}

func main() {

	var c compromisso

	c.dataEspecifica = "19/01/1988"
	c.horaEspecifica = "16:35:44"
	fmt.Println("A data do seu compromisso é", c.dataEspecifica, "e seu horário é", c.horaEspecifica)

}

package main

import "fmt"

type horario struct {
	pHora    *int
	pMinuto  *int
	pSegundo *int
}

func main() {

	var hoje horario

	var hora = 200
	var minuto = 300
	var segundo = 400

	hoje.pHora = &hora
	hoje.pMinuto = &minuto
	hoje.pSegundo = &segundo

	fmt.Println("Hora:", *hoje.pHora)
	fmt.Println("Minuto:", *hoje.pMinuto)
	fmt.Println("Segundo:", *hoje.pSegundo)

	*hoje.pSegundo = 1000
	fmt.Println("Segundo Alterado:", *hoje.pSegundo)
}

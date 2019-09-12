package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 11:
		fallthrough
	// Uma declaração de "fallthrough" transfere o controle para
	//a primeira instrução da próxima cláusula de caso em uma expressão "switch" .
	case 10:
		return "asdfghn"
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {

	fmt.Println(notaParaConceito(10))

	fmt.Println(notaParaConceito(9))

	fmt.Println(notaParaConceito(6.9))

	fmt.Println(notaParaConceito(2.1))

}

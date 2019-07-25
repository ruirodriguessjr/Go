package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 7 && n < 9:
		return "B"
	case n >= 5 && n < 7:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n <= 2:
		return "E"
	default:
		return "Nota Inválida"

	}

}

func main() {
	fmt.Println(notaParaConceito(9.1))
	fmt.Println(notaParaConceito(8.1))


}

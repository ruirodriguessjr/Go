package main

import "fmt"

func main() {

	// var aprovados map[int]string
	// mapas devem ser inicializados
	pessoas := map[string]map[int]float64{
		"Gabriel":      {31: 1.64},
		"Rui":          {32: 1.84},
		"Aline":        {25: 1.70},
		"Guga Pereira": {13: 1.35},
		"José João":    {38: 1.90},
	}

	fmt.Println(pessoas)

	for nome := range pessoas {
		fmt.Println(nome)
	}

}

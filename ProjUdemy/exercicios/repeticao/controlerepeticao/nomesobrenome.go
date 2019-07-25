package main

import "fmt"

func main() {

	nome := "Rui"

	sobrenome := "Rui"

	for i := 0; i < len(nome); i++ {
		if nome[i] == sobrenome[i] {
			fmt.Println("Nomes Iguais")
			break
		} else {
			fmt.Println("Nomes Diferentes")
			break
		}
	}
}

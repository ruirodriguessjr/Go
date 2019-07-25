package main

import "fmt"

func vogal(letra string) string {
	if letra == "a" || letra == "e" || letra == "i" || letra == "o" || letra == "u" {
		return "Sim"
	} else {
		return "Não"
	}

}

func main() {

	letra := "a"
	fmt.Println(vogal(letra))

}

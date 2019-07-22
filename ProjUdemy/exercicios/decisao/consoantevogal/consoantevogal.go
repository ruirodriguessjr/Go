package main

import "fmt"

func main() {

	vogais := []string{"a", "A", "e", "E", "i", "I", "o", "O", "u", "U"}
	letra := "a"

	for _, cont := range vogais {
		if letra != vogais[] {
			fmt.Printf("A letra %s, Não é vogal!", letra)
			break
		}
	
	}
}

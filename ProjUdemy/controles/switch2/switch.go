package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	switch { // Semre que eu não passar condição no meu switch, 
		//ele busca a condição True
	case t.Hour() < 12:
		fmt.Println("Bom dia...")
	case t.Hour() < 18:
		fmt.Println("Boa tarde...")
	default:
		fmt.Println("Boa noite...")
	}

}

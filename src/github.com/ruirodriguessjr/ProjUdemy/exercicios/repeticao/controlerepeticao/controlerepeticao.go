package main

import "fmt"

func main() {

	n := -10

	for {
		if n < 0 || n > 10 {
			fmt.Println("Idiota é entre 0 e 10")
			break
		} else {
			fmt.Println("Número digitado foi ", n)
			break
		}
	}
}

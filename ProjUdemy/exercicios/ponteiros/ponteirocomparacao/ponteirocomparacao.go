package main

import "fmt"

func main() {

	var a string = "Rui"
	var b string = "Victor Viado"

	var EndA *string
	EndA = &a
	var EndB *string
	EndB = &b
	fmt.Println("Valor de A:", *EndA)
	fmt.Println("Endereço de A:", EndA)
	fmt.Println("Valor de B:", *EndB)
	fmt.Println("Endereço de B:", EndB)
	if *EndA > *EndB {
		fmt.Println("Valor de A:", a)
		fmt.Println("Maior Endereço é o de A:", EndA)
	} else {
		fmt.Println("Valor de B:", b)
		fmt.Println("Maior Endereço é o de B:", EndB)
	}
	fmt.Println(EndA, *EndA, EndB, *EndB)

}

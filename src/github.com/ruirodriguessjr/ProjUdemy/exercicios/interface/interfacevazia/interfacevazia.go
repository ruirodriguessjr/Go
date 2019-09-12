package main

import "fmt"

func main() {

	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	// 2º Parametro retorna bool
	s, ok := i.(string)
	fmt.Println(s, ok)

	// 2º Parametro retorna bool
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Caso eu não passe meu 2º parâmetro,
	// Em caso de erro, ele da um panic, deixa de executar meu programa
	c := i.(float64) // Panic
	fmt.Println(c)

	fmt.Println("asxfhgjdbnqosadjk")
}

package main

import "fmt"

func main() {

	// Declarando o tipo de dado do meu slice e a quantidade máxima que ele vai receber
	// Lembrando que não é um limite, é apenas inicial a quantidade
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// Mostrando qual a quantidade  e a capacidade do slice pelo método Make
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

	// Mostrando que a capacidade é aumentada automaticamente quando atingimos seu máximo inicializado
	d := make([]int, 0, 10)
	fmt.Println(len(d), cap(d))
	for i := 0; cap(d) < 100; i++ {
		d = append(d, i)
		fmt.Println(d, len(d), cap(d))
	}

}

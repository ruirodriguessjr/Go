package main

import "fmt"

// Não retorna parâmetros
func f1() {
	fmt.Println("F1")
}

// Recebe dois parâmetros mas não retorna nada
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

// Não recebe parâmetro mas retorna algo
func f3() string {
	return "F3"
}

// Recebe dois parâmetros e retorna um único parâmetro
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

// Não recebe parâmetros, mas retorna 2 valores do tipo prototipado no caso (String)
// Tenho que tratar ambos os retornos, podendo ignorar um dos retornos ultilizando o (_),
// do contrário não é uma boa prática não tratar a quantidade certa de retorno que a função retorna.
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")
	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}

/* Intuito da função é receber parâmetros de outras funções
e trabalhar elas sem alterar o seu externo,
para assim retornar apenas algum tipo de resultado
ou tratamento usando os dados externos que lhe foram de certa forma emprestados. */

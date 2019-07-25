package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	const PI float64 = 3.14
	var raio = 3.2 // Tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma variavel
	area := PI * math.Pow(raio, 2)
	fmt.Println("A area da circunferência é:", area)

	const (
		a = 1
		b = 2
	)
	fmt.Println(a, b)

	var (
		c = 3
		d = 4
	)
	fmt.Println(c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)



	nome := "Rui"
	idade := 31
	versao := 1.1

	fmt.Println("Olá sr.", nome, "sua idade é de", idade, "anos.")
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao))
	// TypeOf para saber o tipo da variável.

	fmt.Println("1 - Iniciar Monitoramento: ")
	fmt.Println("2 - Exibir Logs: ")
	fmt.Println("0 - Sair do Programa: ")

	var comando int
	// fmt.Scanf("%d", &comando)
	// Função Scan, permite que eu apenas passe a variável como parâmetro,
	// pois eu já defini o tipo dela
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variável comando é:", &comando)
	fmt.Println("O comando escolhido foi:", comando)

}

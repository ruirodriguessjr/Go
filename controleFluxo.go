package main

import "fmt"
import "reflect" // Esse pacote utiliza uma função chamada TypeOf.

func main() {

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
	if comando == 1 {
		fmt.Println("Iniciar Monitoramento")
		fmt.Println("Espaço de Memória Ultilizado:", &comando)
	} else if comando == 2 {
		fmt.Println("Exibir Logs")
		fmt.Println("Espaço de Memória Ultilizado:", &comando)
	} else if comando == 0 {
		fmt.Println("Sair do Programa")
	} else {
		fmt.Println("Não conheço esse comando!")
	}

}

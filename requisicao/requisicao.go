package main

import "fmt"
import "reflect"  // Esse pacote utiliza uma função chamada TypeOf.
import "os"       // Retorna para o meu sistema operacional que ele saiu com sucesso
import "net/http" //Pacote com funções prontas, onde faz GET, POST, essa configuração, acesso web

func main() {

	// Quando passo o For sem nenhuma característica dentro dele, ou seja,
	// ele vazio, estou simulando o funcionamento do while True.
	// Ficando dentro de um Loop Infinito.
	for {

		exibeintroducao()
		fmt.Println()
		exibeMenu()
		fmt.Println()
		comando := lerComando()
		fmt.Println()

		// O switch não necessita de brak,
		//ele só vai executar o case escolhido.
		switch comando {
		case 0:
			fmt.Println("Sair do Programa.")
			os.Exit(0)
			//
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibir Logs.")
		default:
			fmt.Println("Este comando não existe.")
			os.Exit(-1)

		}

	}

}

func exibeintroducao() {
	nome := "Rui"
	idade := 31
	versao := 1.1

	fmt.Println("Olá sr.", nome, "sua idade é de", idade, "anos.")
	fmt.Println("Este programa está na versão", versao)
	// TypeOf para saber o tipo da variável.
	fmt.Println("Nome é um(a):", reflect.TypeOf(nome))
	fmt.Println("Idade é um(a):", reflect.TypeOf(idade))
	fmt.Println("Versão é um(a):", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println("1 - Iniciando Monitoramento: ")
	fmt.Println("2 - Exibindo Logs: ")
	fmt.Println("0 - Saindo do Programa: ")
}

func lerComando() int {
	var comandoLido int
	// fmt.Scanf("%d", &comando)
	// Função Scan, permite que eu apenas passe a variável como parâmetro,
	// pois eu já defini o tipo dela
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		fmt.Println()
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		fmt.Println()
	}
}

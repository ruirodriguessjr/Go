package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv" // Conversor de vários tipos para String
	"strings"
	"time"
)

// Esse pacote utiliza uma função chamada TypeOf.
// Retorna para o meu sistema operacional que ele saiu com sucesso
//Pacote com funções prontas, onde faz GET, POST, essa configuração, acesso web
// Pacote para manipulação de tempo.
// Pacote io/ioutil implementa algumas funções de utilitário de Eentrada e Saída de dados

// Contantes são valores imutáveis,
// onde no caso estou usando para deixar meu codigo mais legível,
// onde quem ler o mesmo, possa entender o que está sendo feito
// ao invés de ter que interpretar números dentro do código.
const monitoramentos = 4
const delay = 3

func main() {

	exibeintroducao()
	fmt.Println()

	/*// Passo dois parâmetros com o nome(conteúdo) dele e o status dele
	registraLog("site-falso", false)
	fmt.Println()*/

	lerSitesDoArquivo()
	fmt.Println()

	// Quando passo o For sem nenhuma característica dentro dele, ou seja,
	// ele vazio, estou simulando o funcionamento do while True.
	// Ficando dentro de um Loop Infinito.
	for {

		exibeMenu()
		fmt.Println()

		comando := lerComando()
		fmt.Println()

		// O switch não necessita de brak,
		//ele só vai executar o case escolhido.
		switch comando {
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
			//
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibir Logs...")
			imprimeLogs()
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
	// Slice de sites para otimizar meu código de maneira dinãmica.
	sites := []string{"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	// Controlando quantas vezes que eu quero testar a requisição com meu site.
	for i := 0; i < monitoramentos; i++ {
		// Varrendo meu slice onde:
		// sites vai receber o intervalo(range) dinâmicamente entre o valor inicial de (i),
		// até o valor máximo contido dentro de (sites).
		fmt.Println("Testando a", i+1, "ª vez.")
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		// Tempo de espera para testar o próximo.
		time.Sleep(delay * time.Second)

	}

}

// Função que faz eu testar a requisição com o meu site
// Verificando se o tipo de status code é diferente ou igual a 200.
func testaSite(site string) {
	// Variavel err, seria o tratamento do erro, pois GO não tem Try Catch
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		// Função criada para passar o site que esta monitorando e o status dele naquele momento
		// Como está dentro do tratamento de http com sucesso, quero dizer que ele está ok(True).
		registraLog(site, true)
		fmt.Println()

	} else {
		// Passo o status do codigo de http naquele momento
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		// Nesse caso to passando ele dentro do do tratamento de erro do http e não está ok(False)
		// Passo o status do codigo de http naquele momento
		registraLog(site, false)
		fmt.Println()
	}
}

func lerSitesDoArquivo() []string {
	var sites []string

	// 1ª Opção

	// No caso estamos usando uma funcionalidade do pacote "os",
	// que nos permite abrir o arquivo em questão.
	// No caso aqui ela mostra um ponteiro, como se fosse o local de memória onde o arquivo está apontado
	//arquivo, err := os.Open("sites.txt")

	// 2ª Opção

	// No caso estamos usando como exemplo a funcionalidade do pacote "io/ioutil",
	// que nos permite ter acesso de maneira diferent do arquivo em questão.
	// No caso aqui ela converte o arquivo para um array de bytes,
	// que fica mais fácil fazer a sua conversão.
	//arquivo, err := ioutil.ReadFile("sites.txt")
	// Variavel err, seria o tratamento do erro, pois GO não tem Try Catch

	// 3ª Opção

	// No caso estamos usando como exemplo a funcionalidade do pacote "bufio",
	// que nos permite ler linha a linha, delimitando até onde eu quero ler ele dentro da linha,
	// sendo possível eu usar dentro do pacote funções para ser mais facil a ajuda de Entrada e Saída textual.
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//fmt.Println(string(arquivo))

	leitor := bufio.NewReader(arquivo)
	// Delimitando a leitura para cada linha até a sua quebra, no caso o '\n'.

	for {

		linha, err := leitor.ReadString('\n')
		// Função TrimSpace(parametro) faz com que eu quebre os espaços após.
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		// EOF = erro que acontece quando chega no final do arquivo e ele não encontra nada
		if err == io.EOF {
			break // break caso ele encontrar o final do arquivo ele sai do Loop
		}

	}

	arquivo.Close()
	// Sempre que eu abrir meu arquivo eu fecho ele após termianar minha ação

	return sites

}

// Recebendo como parametros uma string e um booleano
// Vai salvar o site seja offline e o online e o status se tava online ou offline
func registraLog(site string, status bool) {
	// OpenFile abre um arquivo
	// O_RDWR lê e escreve dentro do meu arquivo
	// O_CREATE cria o meu arquivo caso ele nao exista
	// O_APPEND adiciona novo item no meu arquivo, não o sobescrevendo.
	// 0666 é uma permissão para o sistema para escrever ou criar um arquivo
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	// os.Open e os.OpenFile que são consideradas de baixo nível,
	// considerando que ele trabalha mais a nivel do sistema operacional
	// Função ReadFile Automaticamente quando abre o arquivo, após a execução da lógica do código,
	// Já faz o fechamento do arquivo logo após o término da execução do mesmo.
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}

package main

import "fmt"

func main() {
	// Retorno de dois parâmetros na minha função
	nome, idade := devolveNomeEIdade()
	fmt.Println("Meu Nome é", nome, "e minha idade é de", idade, "anos.")

	// o (_) é o perador de identificador em branco
	// faz com que eu ignore uma variavel ou valor.
	_, sobrenome := devolveSexoESobrenome()
	fmt.Println(sobrenome)

}

// Função que retorna dois parâmentros no mesmo retorno de tipos diferentes
// Que poderiam ser de mesmo tipo, lembrando que são respectivamente.
func devolveNomeEIdade() (string, int) {
	nome := "Rui"
	idade := 31
	return nome, idade
}

func devolveSexoESobrenome() (string, string) {
	sexo := "Masculino"
	sobrenome := "Rodrigues"
	return sexo, sobrenome
}

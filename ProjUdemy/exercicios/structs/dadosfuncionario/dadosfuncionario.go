package main

import "fmt"

type pessoa struct {
	Nome           string
	Cpf            string
	Sexo           string
	DataNascimento string
}

type funcionario struct {
	pessoa
	codFunc string
	cargo   string
	salario string
}

func main() {

	//p := pessoa{Nome: "Rui Rodrigues", Cpf: "999.999.999-99", Sexo: "Maculino", DataNascimento: "19/01/1988"}
	f := funcionario{pessoa: pessoa{Nome: "Aline Rodrigues",
		Cpf:            "999.999.999-99",
		Sexo:           "Maculino",
		DataNascimento: "19/01/1988"},

		codFunc: "3318-542",
		cargo:   "Estagiário",
		salario: "1275.00"}

	fmt.Println(f)

}

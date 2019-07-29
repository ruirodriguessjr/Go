package main

import "fmt"

func main() {

	nome := "Rui"
	idade := 1
	sexo := "M"
	estadoCivil := "V"

	for i := 0; i < 1; i++ {
		if len(nome) >= 3 {
			fmt.Println("Seu nome é", nome)
		} else {
			fmt.Println("Nomes menor do que 3 caracteres")
		}
		if idade >= 0 && idade <= 150 {
			fmt.Println("Sua Idade é", idade)
		} else {
			fmt.Println("Idade fora do Intervalo")
		}

		switch {
		case sexo == "f" || sexo == "F":
			fmt.Println("Sexo Feminino")
		case sexo == "m" || sexo == "M":
			fmt.Println("Sexo Masculino")
		default:
			fmt.Println("Sexo Indefinido")
		}

		switch {
		case estadoCivil == "s" || estadoCivil == "S":
			fmt.Println("Solteiro")
		case estadoCivil == "c" || estadoCivil == "C":
			fmt.Println("Casado")
		case estadoCivil == "v" || estadoCivil == "V":
			fmt.Println("Viúvo")
		case estadoCivil == "d" || estadoCivil == "D":
			fmt.Println("Divorciado")
		default:
			fmt.Println("Estado Civil Inválido")
		}

	}
}

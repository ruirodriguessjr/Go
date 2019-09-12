package main

import "fmt"

func main() {

	valorHora := 10.63
	horasDia := valorHora * 6
	horasSemana := horasDia * 5
	horasMes := horasSemana * 4

	fmt.Printf("O total de valor em horas no dia é R$%.2f", horasDia)
	fmt.Printf("\nO total de valor em horas no dia é R$%.2f", horasSemana)
	fmt.Printf("\nO total de valor em horas no dia é R$%.2f", horasMes)

}

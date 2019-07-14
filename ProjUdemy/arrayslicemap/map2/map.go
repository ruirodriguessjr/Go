package main

import "fmt"

func main() {

	// Inicialização de um map 
	// com os tipos correspondentes dos retornos solicitados
	funcsESalarios := map[string]float64{
		"José João": 11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior": 1200.0,
 	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")
	for nome, salario := range funcsESalarios {
		fmt.Println("========================")
		fmt.Println("Nome: ", nome, "\nSalario: ", salario)
 	}
}
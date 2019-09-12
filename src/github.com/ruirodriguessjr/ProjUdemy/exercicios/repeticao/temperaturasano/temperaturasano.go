package main

import "fmt"

func main() {

	temp := map[string]float64{
		"Janeiro":   34.8,
		"Fevereiro": 31.5,
		"Março":     30.9,
		"Abril":     29.5,
		"Maio":      28.4,
		"Junho":     19.4,
		"Julho":     15.3,
		"Agosto":    25.4,
		"Setembro":  27.5,
		"Outubro":   30.4,
		"Novembro":  33.6,
		"Dezembro":  34.4,
	}

	total := 0.0
	for _, t := range temp {
		total += t
	}

	media := total / 12
	fmt.Printf("A média total das temperaturas do ano foi: %.2f", media)
	fmt.Println("\n============================================================")
	for m, t := range temp {
		if t > media {
			fmt.Printf("Temperatura Acima da Média: %s fez %.2fº Celsius.\n", m, t)
		}
	}
	fmt.Println("============================================================")
	for m, t := range temp {
		if t <= media {
			fmt.Printf("Temperatura Abaixo da Média: %s fez %.2fº Celsius.\n", m, t)
		}
	}

}

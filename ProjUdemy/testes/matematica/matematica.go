package matematica

import (
	"fmt"
	"strconv"
)

// Mock = comportamento artificial para uma dependecia que um metodo tenha
// interferencia onde ao invés de chamar a função de acesso a banco ou alguma função que impactaria direta
//no caso crio funções mocks onde ela é criada e chamada apenas para testar unitariamente

// Média é repsonsável por calcular o que você ja sabe
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}

package main

import (
	"fmt"

	"github.com/ruirodriguessjr/area"
	// Pacote area é buscado dentro do caminho de importação de dependência acima
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	fmt.Println(area.Pi)
}

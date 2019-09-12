package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Mostra o número de Processadores que a máquina tem.
	fmt.Println(runtime.NumCPU())
}

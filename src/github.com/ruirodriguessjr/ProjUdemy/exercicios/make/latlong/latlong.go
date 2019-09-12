package main

import "fmt"

type Vertex struct {
	lat, long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
	// make([]int Comprimento, Capacidade)
	numeros := make([]int, 5, 10)
	fmt.Println(numeros[0:2])
}

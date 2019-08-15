package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)

	k := Vertex{1, 2}
	p := &k
	p.X = 1e2
	fmt.Println(k, p, p.X)

	i, j := 42, 2701
	po := &i        // po recebe o endereço de memória de i
	fmt.Println(po) // Imprimindo o endereõ da memória
	fmt.Println(*po)
	*po = 21       // Referenciando ao endereço de memória o valor 21
	fmt.Println(i) // Imprimindo o valor referente ao endereço da memória = 21
	po = &j        // Passando o endereço da memória de J para po
	fmt.Println(po)
	*po = *po / 37
	fmt.Println(j) // 73

}

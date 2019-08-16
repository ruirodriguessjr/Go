package main

import "fmt"

// Ao criar um type direto, sem precisar ser dentro de uma struct
// Automaticamente eu estou dizendo o tipo dele para implementação
// O fato de eu não implementar uma interface nele,
// Não significa que ele não pode ser implementado,
// Basta eu adicionar meu método de interface para ele,
//que ele será implementado pela interface.
type IPAddress [4]byte

func (i IPAddress) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])

}

func main() {
	hosts := map[string]IPAddress{
		"loopback": {127, 0, 0, 1},
		"quad9":    {9, 9, 9, 9},
	}
	for k, v := range hosts {
		fmt.Printf("%v ==> %v\n", k, v)
	}
}

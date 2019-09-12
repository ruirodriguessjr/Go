package main

import "fmt"

type animal struct {
	qtdPes  int
	corPelo string
	som     string
}

type cachorro struct {
	animal
	come  string
	corre bool
}

type gato struct {
	animal
	come  string
	corre bool
}

type peixe struct {
	animal
	come  string
	corre bool
}

func main() {

	//c := cachorro{animal: animal{qtdPes: 4, corPelo: "Marrom", som: "AUAU"}, come: "Comida", corre: true}
	c := cachorro{}
	c.come = "bosta"
	g := gato{animal: animal{qtdPes: 4, corPelo: "Cinza", som: "Miau"}, come: "Ração", corre: true}
	p := peixe{animal: animal{qtdPes: 0, corPelo: "Cinza", som: "Glob"}, come: "Corais", corre: false}

	fmt.Println(c)
	fmt.Println(g)
	fmt.Println(p)
}

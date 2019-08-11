package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos caso queira

}

type bwm7 struct{}

func (b bwm7) ligarTurbo() {

	fmt.Println("Turbo...")

}

func (b bwm7) fazerBaliza() {

	fmt.Println("Baliza...")

}

func main() {

	// Referencia dos métodos de interface,
	// podem ser acessadas individualmente ou por meio de interface
	// que ultiliza as duas interfaces(esportivo, luxuoso) dentro dela.
	var b esportivoLuxuoso = bwm7{}
	//var b esportivo = bwm7{}
	//var b luxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()

}

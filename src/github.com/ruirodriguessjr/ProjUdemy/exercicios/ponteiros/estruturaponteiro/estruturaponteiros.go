package main

import "fmt"

type horario struct {
	hora    int
	minuto  int
	segundo int
}

func main() {

	var agora horario
	var antes horario
	var depois *horario
	depois = &agora

	soma := 100

	(*depois).hora = 20
	(*depois).minuto = 80
	(*depois).segundo = 50

	antes.hora = soma + depois.segundo
	antes.minuto = agora.hora + depois.minuto
	antes.segundo = depois.minuto + depois.segundo

	fmt.Println(antes.hora, ":", antes.minuto, ":", antes.segundo)

}

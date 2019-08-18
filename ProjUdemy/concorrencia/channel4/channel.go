package main

import (
	"fmt"
	"time"
)

var s = 0

func soma(x []int, somatoria chan<- int) {
	for _, v := range x {
		s += v
	}
	somatoria <- s //Enviando dados para o canal

}

func main() {

	vs := []int{7, 2, 8, -9, 4, 0}
	somatoria := make(chan int)
	go soma(vs[:len(vs)/2], somatoria)
	go soma(vs[len(vs)/2:], somatoria)

	x, y := <-somatoria, <-somatoria // recebendo os dados do canal

	time.Sleep(time.Second)

	fmt.Println("Lista completa antes da soma é: ", vs)
	fmt.Println("A somatória dos valores1 é: ", x)
	fmt.Println("A somatória dos valores2 é: ", y)
	fmt.Println("End Function to goroutine")
}

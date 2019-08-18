package main

import (
	"fmt"
	"time"
)

var s = 0

func soma(x []int) {
	for _, v := range x {
		s += v
	}

}

func main() {

	vs := []int{7, 2, 8, -9, 4, 0}
	go soma(vs[:len(vs)/2])
	go soma(vs[len(vs)/2:])
	time.Sleep(250 * time.Millisecond)
	fmt.Println("Lista completa antes da soma é: ", vs)
	fmt.Println("A somatória dos valores é: ", s)
	fmt.Println("End Function to goroutine")
}

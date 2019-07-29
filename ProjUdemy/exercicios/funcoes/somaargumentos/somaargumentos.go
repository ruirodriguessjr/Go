package main

import "fmt"

func soma(a int, b int, c int) int{
	return a + b + c
}

func main(){
	a := 1
	b := 2
	c := 3
	fmt.Println(soma(a, b, c))
}
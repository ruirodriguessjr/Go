package main

import "fmt"

func main() {

	a := 10
	b := 1.0
	c := "10"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var pa *int = &a
	*pa = 20
	var pb *float64 = &b
	*pb = 20.00
	var pc *string = &c
	*pc = "Rui"
	fmt.Println(*pa)
	fmt.Println(*pb)
	fmt.Println(*pc)

}

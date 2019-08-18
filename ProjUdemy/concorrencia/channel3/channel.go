package main

import "fmt"

func multiply(num int, out chan<- int) {
	result := num * 2
	out <- result
}

func main() {
	n := 3
	out := make(chan int)
	go multiply(n, out)
	fmt.Println(<-out)
}

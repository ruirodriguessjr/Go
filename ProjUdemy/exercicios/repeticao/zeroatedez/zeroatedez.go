package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		for j := 10; j >= 0; j-- {
			fmt.Println(i, j)
		}
	}
}

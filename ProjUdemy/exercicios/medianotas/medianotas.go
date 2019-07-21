package main

import "fmt"

func media(m float64) float64{
	return m / 4
}

func main() {

	nums := []float64{8.9, 6.7, 5.8, 7.8}
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("soma:", sum)

	m := media(sum)

	fmt.Println(m)

}

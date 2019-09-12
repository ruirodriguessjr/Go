package main

import "fmt"

func soma(n [5]int) int {
	var s int
	for i := 0; i < len(n)-1; i++ {
		s += n[i]
	}
	return s
}

func mult(n [5]int) int {
	var m int
	for i := 1; i < len(n); i++ {
		m *= n[i]
		fmt.Println(m)
	}
	return m
}

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	var s, m int
	s = soma(nums)
	m = mult(nums)
	fmt.Println("A soma dos números é:", s)
	fmt.Println("A multiplicação dos números é:", m)

}

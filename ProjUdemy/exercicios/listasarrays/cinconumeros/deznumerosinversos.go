package main

import "fmt"

func main() {

	array := [10]float64{1.2, 2.4, 2.1, 5.4, 6.7, 7.8, 9.1, 10.3, 3.7, 4.7}

	for i := 9; i >= 0; i-- {
		fmt.Println(array[i])
	}

}

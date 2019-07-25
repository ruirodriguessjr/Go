package main

import "fmt"

func main() {

	array := [10]string{"a", "e", "i", "o", "s", "d", "f", "g", "h", "j"}
	var consoante [10]string
	for i := 0; i < len(array); i++ {
		if array[i] != "a" && array[i] != "e" && array[i] != "i" && array[i] != "o" && array[i] != "u" {
			consoante[i] = array[i]
		}
	}
	fmt.Println(consoante)

}

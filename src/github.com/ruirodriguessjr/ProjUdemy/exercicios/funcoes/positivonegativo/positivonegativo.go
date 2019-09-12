package main

import "fmt"

func posNeg(n int) string {
	if n > 0 {
		return "Positivo"
	} else {
		return "Negativo"
	}
}
func main() {
	n := -1
	m := 10

	fmt.Println(posNeg(n))
	fmt.Println(posNeg(m))

}

package main

import (
	"fmt"
	"strings"
)

func main(){
	m := make(map[string]int)

	str := `I ate a donut. Then I ate another donut.`

	for _, w := range strings.Fields(str){
		m[w]++
	}

	fmt.Printf("Contador de palavras: %v", m)
}
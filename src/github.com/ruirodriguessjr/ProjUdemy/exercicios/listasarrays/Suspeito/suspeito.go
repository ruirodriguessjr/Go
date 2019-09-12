package main

import (
	"fmt"
	"strconv"
)

func main() {

	questions := map[string]bool{
		"Telefonou para a vítima?":   true,
		"Esteve no local do crime?":  true,
		"Mora perto da vítima?":      true,
		"Devia para a vítima?":       true,
		"Já trabalhou com a vítima?": true,
	}

	countTrue, contrFalse := 0, 0

	for m, v := range questions {
		fmt.Printf("%s %s\n", m, strconv.FormatBool(v))
		if v == true {
			countTrue++
		} else {
			contrFalse++
		}
	}
	if countTrue < 2 {
		fmt.Println("Inocente!")
	} else if countTrue == 2 {
		fmt.Println("Suspeita!")
	} else if countTrue == 3 || countTrue == 4 {
		fmt.Println("Cúmplice!")
	} else {
		fmt.Println("Assassino!")
	}

}

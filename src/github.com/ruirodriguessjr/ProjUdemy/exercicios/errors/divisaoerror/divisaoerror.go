package main

import "fmt"

type error interface {
	Error() string
}

type ErrDivNegativo int

func (e ErrDivNegativo) Error() string {
	return fmt.Sprintf("O  divisor não pode ser Negativo ou Zero: %d", int(e))
}

func Divisao(v1, v2 int) (float64, error) {
	if v2 <= 0 {
		return 0, ErrDivNegativo(v2)
	}
	r := float64(v1 / v2)
	return r, nil
}

func main() {
	fmt.Println(Divisao(8, 2))
	fmt.Println(Divisao(8, 0))
	fmt.Println(Divisao(8, -4))

}

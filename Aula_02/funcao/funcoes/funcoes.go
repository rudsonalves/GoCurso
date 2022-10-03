package main

import "fmt"

func divisao(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("y deve ser diferente de zero")
	}
	return x / y, nil
}

func main() {
	v, ok := divisao(3., 0.)
	fmt.Println(v, ok)

	v, ok = divisao(3., 2.)
	fmt.Println(v, ok)
}

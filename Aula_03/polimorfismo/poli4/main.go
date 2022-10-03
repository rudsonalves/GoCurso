package main

import "fmt"

func main() {
	v1 := Vetor{3, 5}
	v2 := Vetor{2, 7}
	fmt.Println("v1:", v1, " v2:", v2)
	v3 := v1.Soma(v2)
	fmt.Println("v3:", v3)
	v4 := v1.Produto(2)
	fmt.Println("v4:", v4)
	v5 := v1.Produto(3.)
	fmt.Println("v5:", v5)
	v6 := v1.Produto(v2)
	fmt.Println("v6:", v6)
}

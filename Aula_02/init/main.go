package main

import "fmt"

var start = 0

func init() {
	start = 10
	fmt.Println("Este é o init() em main.go...")
}

func main() {
	fmt.Println("Está é a main()...")
	fmt.Println("Start:", start)
	DigaOla("Sandra")
}

func init() {
	fmt.Println("Está é uma redeclaração de init no mesmo arquivo.go...")
}

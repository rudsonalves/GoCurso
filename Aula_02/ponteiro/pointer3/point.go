package main

import "fmt"

func Dobrar(x *int) {
	*x *= 2
	fmt.Println(x)
}

func main() {
	a := 7
	Dobrar(&a)

	fmt.Println(a)
}

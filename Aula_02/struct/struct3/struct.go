package main

import "fmt"

type Ponto struct {
	x int
	y int
}

func main() {
	a := Ponto{2, 3}
	b := a
	b.x = 5
	fmt.Println(a, b)
}

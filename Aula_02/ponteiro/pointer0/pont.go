package main

import "fmt"

func soma(x, y int, z *int) {
	*z = x + y
}

func main() {
	a := 5
	var b *int = &a
	*b += 2

	fmt.Println(a, b)

	x, y, z := 3, 5, 1

	soma(x, y, &z)

	fmt.Println(x, y, z)
}

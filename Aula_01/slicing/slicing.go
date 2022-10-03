package main

import "fmt"

func main() {
	x := [6]int{1, 2, 3, 4, 5, 6}
	y := make([]int, 2)
	copy(y, x[2:])
	fmt.Println(y)
	y[1] = 0
	fmt.Println(x, y)
	var z []int
	fmt.Println(z, len(z))
	z = append(z, 1)
	fmt.Println(z, len(z), cap(z))
}

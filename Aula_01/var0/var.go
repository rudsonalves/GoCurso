package main

import (
	"fmt"
)

func main() {
	// var (
	// 	x float64 = 2
	// 	y int64
	// )

	// fmt.Printf("x = %.2f   y = %d\n", x, y)

	// w := int8(5)
	// z := 3 + 8i
	// v := 12.
	// fmt.Printf("w: %v (%[1]T)\n", w)
	// fmt.Printf("z: %v (%[1]T)\n", z)
	// fmt.Printf("v: %v (%[1]T)\n", v)

	i, j := 12, 5.
	fmt.Printf("i: %v (%[1]T)\n", i)
	fmt.Printf("j: %v (%[1]T)\n", j)

	var (
		a, b int = 5, 6
	)
	fmt.Printf("a: %v (%[1]T)\n", a)
	fmt.Printf("b: %v (%[1]T)\n", b)

	i = a + b
	fmt.Printf("i: %v (%[1]T)\n", i)

	a, c := 4, 7
	fmt.Println(a, c)

	a, d := "1.5", 2
	fmt.Println(a, d)

	/*
		const pi = 3.1515
	*/

	// const (
	// 	_ = 1 << (10 * iota)
	// 	KiB
	// 	MiB
	// 	GiB
	// 	TiB
	// )

	// fmt.Println(KiB, int(math.Pow(2, 10)))
	// fmt.Println(MiB, int(math.Pow(2, 20)))
	// fmt.Println(GiB, int(math.Pow(2, 30)))
	// fmt.Println(TiB, int(math.Pow(2, 40)))
}

package main

import "fmt"

func main() {
	fat := func(n int) int {
		f := 1
		for i := n; i > 1; i-- {
			f *= i
		}
		return f
	}

	v := fat(5)
	fmt.Println(v)
}

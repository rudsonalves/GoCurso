package main

import "fmt"

func main() {
	var a interface{}

	fmt.Printf("a: %v (tipo %[1]T)\n", a)
	a = 5
	fmt.Printf("a: %v (tipo %[1]T)\n", a)
	a = "rudson"
	fmt.Printf("a: %v (tipo %[1]T)\n", a)
	a = 2 + 3i
	fmt.Printf("a: %v (tipo %[1]T)\n", a)
}

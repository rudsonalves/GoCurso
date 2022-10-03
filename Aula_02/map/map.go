package main

import "fmt"

func main() {
	var t = make(map[string]int)

	t["abobora"] = 5

	fmt.Printf("t: %v (%[1]T)", t)
}

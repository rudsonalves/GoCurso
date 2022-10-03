package main

import (
	"fmt"
)

func main() {
	var pop = make(map[string]float64)
	fmt.Println(pop == nil)
	pop["Brasil"] = 213317639
	pop["China"] = 1411780000
	fmt.Println(pop)
}

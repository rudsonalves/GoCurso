package main

import (
	"fmt"
)

func main() {
	n := 99

	for {
		if EPrimo(n) {
			n--
			continue
		}
		if n < 1 {
			break
		} else {
			fmt.Printf("%2d ", n)
		}

		n--
	}
	fmt.Println("")
}

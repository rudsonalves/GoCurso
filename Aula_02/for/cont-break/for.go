package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			continue
		}
		if i%31 == 0 {
			break
		}
		fmt.Printf("%2d ", i)
	}
	fmt.Println("\nFim")
}

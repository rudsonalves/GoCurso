package main

import "fmt"

func main() {
	fmt.Print("   ")
	for d := 1; d <= 31; d++ {
		if d%7 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%2d ", d)
	}
}

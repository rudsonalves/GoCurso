package main

import "fmt"

func main() {
loop_i:
	for i := 0; i < 10; i++ {
	loop_j:
		for j := 0; j < 10; j++ {
			if i%2 == 0 {
				continue loop_i
			} else if j%3 == 0 {
				continue loop_j
			} else if i+j > 17 {
				break loop_i
			}
			fmt.Printf("(%d, %d) ", i, j)
		}
	}
	fmt.Println("")
}

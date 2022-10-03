package main

import "fmt"

func print(id int, ch <-chan int) {
	fmt.Printf("(%d) channel: %d\n", id, <-ch)
}

func main() {
	ch := make(chan int)

	go print(1, ch)
	go print(2, ch)
	go print(3, ch)

	ch <- 0
	ch <- 1
	ch <- 2
	ch <- 3
}

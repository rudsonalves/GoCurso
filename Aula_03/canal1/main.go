package main

import "fmt"

func print(id int, ch chan<- int) {
	ch <- id
}

func main() {
	ch := make(chan int)

	go print(1, ch)
	go print(2, ch)
	go print(3, ch)

	fmt.Printf("Respondeu id: %d\n", <-ch)
	fmt.Printf("Respondeu id: %d\n", <-ch)
	fmt.Printf("Respondeu id: %d\n", <-ch)
}

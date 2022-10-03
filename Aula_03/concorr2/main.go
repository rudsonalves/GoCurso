package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func print(ch <-chan int) {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		fmt.Printf("Print %02d (len: %d, loop: %d)\n", <-ch, len(ch), i)
	}
}

func main() {
	ch := make(chan int, 3)

	wg.Add(1)
	go print(ch)

	for i := 0; i < 6; i++ {
		value := i * 10
		fmt.Printf("Write %02d\n", value)
		ch <- value
	}
	fmt.Println("Wait...")
	wg.Wait()
}

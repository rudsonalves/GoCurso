package main

import (
	"fmt"
	"time"
)

func contador(id int, t time.Duration) {
	for i := 0; i < 10; i++ {
		time.Sleep(t)
		fmt.Printf("Id %02d: Step: %d\n", id, i)
	}
}

func main() {
	fmt.Println("Início...")

	go contador(1, time.Second)
	go contador(2, time.Second)
	go contador(3, time.Second)

	time.Sleep(time.Second * 5)
	fmt.Println("Fim...")
}

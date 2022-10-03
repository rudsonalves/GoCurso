package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func contador(id int, t time.Duration) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(t)
		fmt.Printf("Id %02d: Step: %d\n", id, i)
	}
}

func main() {
	fmt.Println("Início...")

	rand.Seed(time.Now().UnixMicro())

	for id := 0; id < 3; id++ {
		wg.Add(1)
		t := time.Duration(100 + rand.Intn(100))
		go contador(id, time.Microsecond*t)
	}

	wg.Wait()
	fmt.Println("Fim...")
}

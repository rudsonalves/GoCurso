package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func connectDB(server string, ch chan<- string, done <-chan struct{}) {
	defer wg.Done()
	fmt.Printf("Try connection in %s...\n", server)
	t := time.Duration(20 + rand.Int63n(10))
	time.Sleep(time.Microsecond * t)
	select {
	case ch <- server:
		return
	case <-done:
		fmt.Printf("Server %s close!\n", server)
		return
	}
}

func main() {
	serverList := []string{"Newton", "Einstein", "Plank", "Fermi"}
	ch := make(chan string)
	done := make(chan struct{}, len(serverList)-1)

	rand.Seed(time.Now().UnixNano())

	for _, s := range serverList {
		wg.Add(1)
		go connectDB(s, ch, done)
	}

	doneAll := func() {
		for i := 0; i < cap(done); i++ {
			done <- struct{}{}
		}
		close(done)
	}

	server := <-ch
	doneAll()

	fmt.Printf("Server %s CONNECTED.\n", server)
	wg.Wait()
	close(ch)
}

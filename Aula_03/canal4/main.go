package main

import (
	"fmt"
	"math/rand"
	"time"
)

func connectDB(server string, ch chan<- string) {
	fmt.Printf("Try connection in %s...\n", server)
	t := time.Duration(20 + rand.Int63n(10))
	time.Sleep(time.Microsecond * t)
	ch <- fmt.Sprintf("%s connected!", server)
}

func main() {
	serverList := []string{"Newton", "Einstein", "Plank", "Fermi"}
	var c []chan string

	rand.Seed(time.Now().UnixNano())

	for i, server := range serverList {
		c = append(c, make(chan string))
		go connectDB(server, c[i])
	}

	select {
	case s := <-c[0]:
		fmt.Println(s)
	case s := <-c[1]:
		fmt.Println(s)
	case s := <-c[2]:
		fmt.Println(s)
	case s := <-c[3]:
		fmt.Println(s)
	}
}

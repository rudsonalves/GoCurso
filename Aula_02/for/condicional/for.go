package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())

	d := rand.Intn(10)

	for d < 8 {
		fmt.Println("d", d)
		d++
	}
}

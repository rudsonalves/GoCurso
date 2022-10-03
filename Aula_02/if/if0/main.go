package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())

	n := rand.Intn(20)
	if n == 0 {
		fmt.Println("n é nulo.")
	} else if n%2 == 0 {
		fmt.Println("n é múltiplo de 2:", n)
	} else if n%3 == 0 {
		fmt.Println("n é múltiplo de 3:", n)
	} else {
		fmt.Printf("n (%d) não é múltiplo de 2 ou 3.", n)
	}
}

package main

import (
	"log"
)

var Primos []int = []int{2, 3, 5, 7, 11}

const MAXPRIMO = 11 * 11

func EPrimo(n int) bool {
	if n > MAXPRIMO {
		log.Fatalf("n deve ser menos que %d", MAXPRIMO)
	}

	for _, p := range Primos {
		if n == p {
			return true
		} else if n%p == 0 {
			return false
		}
	}
	return true
}

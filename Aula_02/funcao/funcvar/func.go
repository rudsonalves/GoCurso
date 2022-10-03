package main

import "fmt"

func soma(nums ...int) int {
	s := 0
	if len(nums) < 1 {
		return s
	}

	for _, valor := range nums {
		s += valor
	}

	return s
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

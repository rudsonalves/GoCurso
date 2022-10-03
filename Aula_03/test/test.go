package main

import "fmt"

func main() {
	a := "0123456789"
	r := []rune(a)
	fmt.Println(r, a)

	words := map[string]int{}

	fmt.Println(words["test"])
}

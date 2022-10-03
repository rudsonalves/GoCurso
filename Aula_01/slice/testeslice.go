package main

import "fmt"

func main() {
	w := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("w: ", w, len(w), cap(w))
	w = append(w, 9, 10)
	fmt.Println("w: ", w, len(w), cap(w))
}

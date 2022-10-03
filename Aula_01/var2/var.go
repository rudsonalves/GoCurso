package main

import "fmt"

func main() {
	w := []int{1, 2, 3, 4}
	w = append(w, 5, 6)
	fmt.Println(w)
	y := []int{7, 8, 9, 0}
	w = append(w, y[:2]...)
	fmt.Println(w)
}

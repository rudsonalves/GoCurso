package main

import "fmt"

func main() {
	var s [4]int
	fmt.Println(s)
	var v [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(v)
	var z = [10]int{1, 1, 5: 4, 8: 77}
	fmt.Println(z)
	var m = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(m)
}

package main

import "fmt"

var a int = 5

func init() {
	fmt.Println("Em init():", a)
}

func main() {
	fmt.Println("No início do main():", a)
	a := 10

	func() {
		a := "15"
		fmt.Println("Na func anônima:", a)
	}()

	if true {
		a := 20
		fmt.Println("No if", a)
	}

	fmt.Println("Em main():", a)
}

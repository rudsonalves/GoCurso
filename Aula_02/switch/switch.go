package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		switch i {
		case 2, 4, 6, 8, 10:
			fmt.Println("É par:", i)
		default:
			fmt.Println("É impar:", i)
		}
	}

	for i := 1; i < 10; i++ {
		switch a := 3 * i; {
		case a%2 == 0:
			fmt.Println("3*i é par:", a)
		case i%3 == 0:
			fmt.Println("I é múltiplo de 3:", i)
		default:
			fmt.Println("3*i impar:", a)
		}
	}
}

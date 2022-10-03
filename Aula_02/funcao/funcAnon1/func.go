package main

import "fmt"

func main() {
	func(nome string) {
		fmt.Printf("Olá %s.\n", nome)
	}("Rudson")
}

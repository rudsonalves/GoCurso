package main

import "fmt"

func main() {
	hello := func(nome string) {
		fmt.Printf("Olá %s\n", nome)
	}

	hello("Gustavo")

}

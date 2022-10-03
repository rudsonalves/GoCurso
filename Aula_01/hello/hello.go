package main

import (
	"fmt"
	"log"
)

func main() {
	var nome, sobrenome string
	fmt.Println("Entre com o seu nome e sobrenome:")
	_, err := fmt.Scan(&nome, &sobrenome)
	if err != nil {
		log.Fatal(err)
	}

	hello(nome, sobrenome)
}

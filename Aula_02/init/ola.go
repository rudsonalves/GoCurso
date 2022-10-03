package main

import "fmt"

func DigaOla(nome string) {
	fmt.Println("Olá", nome, "...")
}

func init() {
	fmt.Println("Está é a init() em Ola.go...")
}

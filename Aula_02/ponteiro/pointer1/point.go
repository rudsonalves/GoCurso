package main

import (
	"fmt"
)

func main() {
	a := new(int)
	*a = 5

	p := new(struct{ nome, sobrenome string })
	p.nome = "Rudson"
	p.sobrenome = "Alves"

	fmt.Println(*a, *p)
	fmt.Println(a, p)
}

package main

import (
	"fmt"
	"strings"
)

func Double(x *int) {
	*x *= 2
}

func main() {
	a := 7
	Double(&a)
	fmt.Println(a)

	p := &struct{ nome, sobrenome string }{"Rudson", "Alves"}
	Capt(p)
	fmt.Println(*p)
}

func Capt(p *struct{ nome, sobrenome string }) {
	p.nome = strings.ToUpper(p.nome)
	p.sobrenome = strings.ToUpper(p.sobrenome)
}

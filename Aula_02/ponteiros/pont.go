package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	Nome      string
	SobreNome string
}

func (p *Pessoa) Capt() {
	p.Nome = strings.ToUpper(p.Nome)
	p.SobreNome = strings.ToUpper(p.SobreNome)
}

func Capt(s *string) {
	*s = strings.ToUpper(*s)
}

func main() {
	a := new(Pessoa)
	a.Nome = "Rudson"
	a.SobreNome = "Alves"
	a.Capt()

	fmt.Println("Valor:", *a, "Ponteiro:", a)

	// b := Pessoa{"Albert", "Einstein"}
	// c := &b

	d := "Albert Einstein"
	Capt(&d)

	fmt.Println(d)
}

package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	nome      string
	sobrenome string
}

func (p *Pessoa) Capt() {
	p.nome = strings.ToUpper(p.nome)
	p.sobrenome = strings.ToUpper(p.sobrenome)
}

func DeCapt(p *Pessoa) {
	p.nome = strings.ToLower(p.nome)
	p.sobrenome = strings.ToUpper(p.sobrenome)
}

func main() {
	p := &Pessoa{"Albert", "Santos"}
	p.Capt()

	fmt.Print(*p)

	DeCapt(p)
	fmt.Println(*p)
}

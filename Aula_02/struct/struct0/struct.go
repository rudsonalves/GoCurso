package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
	Peso  int
}

type PessoaFisica struct {
	Pessoa
	CPF int
}

func (p Pessoa) String() string {
	strf := "Nome: %s, Idade: %d, Peso: %d"
	return fmt.Sprintf(strf, p.Nome, p.Idade, p.Peso)
}

func (p PessoaFisica) String() string {
	strf := "Nome: %s\nIdade: %d\nPeso: %d\nCPF: %d"
	return fmt.Sprintf(strf, p.Nome, p.Idade, p.Peso, p.CPF)
}

func main() {
	p := Pessoa{
		Nome:  "Robert Santos",
		Idade: 36,
		Peso:  83,
	}

	pf := PessoaFisica{
		Pessoa: Pessoa{
			Nome:  "João Alves",
			Idade: 45,
			Peso:  78,
		},
		CPF: 82734617449,
	}

	fmt.Println(p)
	fmt.Printf("\n%v\n", pf)
}

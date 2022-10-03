package main

import "fmt"

type Pai struct {
	Nome  string
	Idade int
	CPF   string
}

func (p Pai) String() string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Cpf: %s", p.Nome, p.Idade, p.CPF)
}

type Filho struct {
	Pai
	Email string
}

func (f Filho) String() string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Cpf: %s, E-mail: %s", f.Nome, f.Idade, f.CPF, f.Email)
}

type Stringer interface {
	String() string
}

func Println(a ...Stringer) {
	for _, v := range a {
		fmt.Println(v.String())
	}
}

func main() {
	p := Pai{"João Walmiro Alves", 85, "123.123.123-73"}
	f := Filho{Pai{"Rudson Alves", 55, "321.321.321-32"}, "rudsonalves67@gmail.com"}

	Println(p)
	Println(f)
}

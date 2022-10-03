package main

import (
	"fmt"
)

type Retangulo struct {
	Base   float64
	Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Base
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Altura + r.Base)
}

func (r *Retangulo) SetAltura(valor float64) {
	r.Altura = valor
}

func (r *Retangulo) SetBase(valor float64) {
	r.Base = valor
}

func (r *Retangulo) SetAll(base, altura float64) {
	r.Base = base
	r.Altura = altura
}

func (r Retangulo) String() string {
	return fmt.Sprintf("Retângulo:\n Base: %.2f\n Altura: %.2f", r.Base, r.Altura)
}

func main() {
	a := Retangulo{
		Base:   2.,
		Altura: 3.,
	}

	a.SetAll(3, 5)

	fmt.Println(a)

	a.Base = 20
	a.Altura = 30

	fmt.Println(a)
}

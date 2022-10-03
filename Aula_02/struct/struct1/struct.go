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

func main() {
	a := Retangulo{
		Base:   2.,
		Altura: 3.,
	}

	fmt.Printf("Retângulo:\n Altura: %.2f\n Base: %.2f\n Área: %.2f\n Perímetro: %.2f\n", a.Altura, a.Base, a.Area(), a.Perimetro())
}

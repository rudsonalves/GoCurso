package main

import (
	"fmt"
	"math"
)

const Pi = math.Pi

type Circulo struct {
	Raio int
}

func (c Circulo) Area() int {
	return int(Pi * float64(c.Raio*c.Raio))
}

func (c Circulo) Perimetro() int {
	return int(2 * Pi * float64(c.Raio))
}

type Retangulo struct {
	Base   int
	Altura int
}

func (r Retangulo) Area() int {
	return r.Base * r.Altura
}

func (r Retangulo) Perimetro() int {
	return 2 * (r.Base + r.Altura)
}

type Quadrado struct {
	Aresta int
}

func (q Quadrado) Area() int {
	return q.Aresta * q.Aresta
}

func (q Quadrado) Perimetro() int {
	return 4 * q.Aresta
}

type Forma interface {
	Area() int
	Perimetro() int
}

func AreaTotal(obj ...Forma) int {
	soma := 0
	for _, f := range obj {
		soma += f.Area()
	}
	return soma
}

func PerimetroTotal(obj ...Forma) int {
	soma := 0
	for _, f := range obj {
		soma += f.Perimetro()
	}
	return soma
}

func main() {
	a := Circulo{23}
	b := Retangulo{12, 32}
	c := Quadrado{25}

	fmt.Println("Área Total:", a.Area()+b.Area()+c.Area())
	fmt.Println("Área Total:", AreaTotal(a, b, c))
	fmt.Println("Perímetro Total:", a.Perimetro()+b.Perimetro()+c.Perimetro())
	fmt.Println("Perímetro Total:", PerimetroTotal(a, b, c))
}

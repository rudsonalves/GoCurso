package main

import (
	"fmt"
	"math"
)

const Pi = math.Pi

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

func (r Retangulo) String() string {
	return fmt.Sprintf("Retângulo:\n  Altura: %.2f\n  Base: %.2f", r.Altura, r.Base)
}

type Quadrado struct {
	Aresta float64
}

func (q Quadrado) Area() float64 {
	return q.Aresta * q.Aresta
}

func (q Quadrado) Perimetro() float64 {
	return 4 * q.Aresta
}

func (q Quadrado) String() string {
	return fmt.Sprintf("Quadrado:\n  Aresta: %.2f", q.Aresta)
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return c.Raio * c.Raio * Pi
}

func (c Circulo) Perimetro() float64 {
	return 2 * Pi * c.Raio
}

func (c Circulo) String() string {
	return fmt.Sprintf("Círculo:\n  Raio: %.2f", c.Raio)
}

type Shaper interface {
	Area() float64
	Perimetro() float64
}

func AreaTotal(shapes ...Shaper) float64 {
	soma := 0.
	for _, s := range shapes {
		soma += s.Area()
	}

	return soma
}

func PerimetroTotal(shapes ...Shaper) float64 {
	soma := 0.
	for _, s := range shapes {
		soma += s.Perimetro()
	}

	return soma
}

func main() {
	r := Retangulo{10, 20}
	q := Quadrado{24}
	c := Circulo{12.5}

	area := AreaTotal(r, q, c)
	perimetro := PerimetroTotal(r, c, q)
	fmt.Println("Area Total:", area)
	fmt.Println("Perímetro Total:", perimetro)
	fmt.Println(c)
	fmt.Println(r)
	fmt.Println(q)
}

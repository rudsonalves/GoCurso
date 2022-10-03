package main

import (
	"fmt"
	"math"
)

type any = interface{}

type Medida struct {
	Valor     float64
	Incerteza float64
}

func (c Medida) String() string {
	return fmt.Sprintf("%.2f ± 𝚫%.2f", c.Valor, c.Incerteza)
}

func (c Medida) Produto(valor any) Medida {
	r := Medida{0., 0.}
	switch v := valor.(type) {
	case int:
		r.Valor = c.Valor * float64(v)
		r.Incerteza = c.Incerteza * float64(v)
	case int64:
		r.Valor = c.Valor * float64(v)
		r.Incerteza = c.Incerteza * float64(v)
	case float32:
		r.Valor = c.Valor * float64(v)
		r.Incerteza = r.Incerteza * float64(v)
	case float64:
		r.Valor = c.Valor * v
		r.Incerteza = r.Incerteza * v
	case Medida:
		r.Valor = c.Valor * v.Valor
		r.Incerteza = math.Abs(c.Valor*v.Incerteza + c.Incerteza*v.Valor)
	default:
		fmt.Printf("Error %v não suportado!", v)
	}
	return r
}

// main function
func main() {
	a := Medida{8., .5}
	b := Medida{1, .2}
	fmt.Println("   a:", a)
	fmt.Println("   b:", b)
	fmt.Println(" 2*a:", a.Produto(2))
	fmt.Println(" a*b:", a.Produto(b))
	fmt.Println("a*as:", a.Produto("as"))
}

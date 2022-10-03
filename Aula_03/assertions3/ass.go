package main

import (
	"log"
)

type Vetor struct {
	X, Y float64
}

func (v Vetor) Soma(a Vetor) Vetor {
	return Vetor{v.X + a.X, v.Y + a.Y}
}

func (v Vetor) Produto(valor interface{}) Vetor {
	b := Vetor{}
	if x, ok := valor.(int); ok {
		b.X = v.X * float64(x)
		b.Y = v.Y * float64(x)
	} else if x, ok := valor.(float64); ok {
		b.X = v.X * x
		b.Y = v.Y * x
	} else {
		log.Fatalf("tipo '%T' não suportado", x)
	}
	return b
}

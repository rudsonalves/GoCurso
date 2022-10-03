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
	switch t := valor.(type) {
	case int:
		x, _ := valor.(int)
		b.X = v.X * float64(x)
		b.Y = v.Y * float64(x)
	case float64:
		x, _ := valor.(float64)
		b.X = v.X * x
		b.Y = v.Y * x
	default:
		log.Fatalf("tipo '%T' não suportado", t)
	}
	return b
}

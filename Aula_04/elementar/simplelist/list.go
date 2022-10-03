package simplelist

import (
	"fmt"
)

type List []float64

func (l *List) Append(e ...float64) {
	*l = append(*l, e...)
}

func (l List) Add(k List) (List, error) {
	if len(l) != len(k) {
		err := fmt.Errorf("o númeor de elementos de l e k são diferentes")
		return List{}, err
	}

	out := List{}
	for id, v := range k {
		out.Append(v + l[id])
	}

	return out, nil
}

func NewList(e ...float64) List {
	out := List{}
	out.Append(e...)

	return out
}

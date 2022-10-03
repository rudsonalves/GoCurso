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
		err := fmt.Errorf("as listas devem ter o mesmo tamanhos para somar")
		return List{}, err
	}

	out := List{}
	for i, v := range l {
		out = append(out, v+k[i])
	}

	return out, nil
}

func (l List) String() string {
	str := "List["
	for _, v := range l {
		str += fmt.Sprintf("%.0f ", v)
	}
	str = str[:len(str)-1] + "]"
	return str
}

func NewList(e ...float64) List {
	l := List{}

	l.Append(e...)
	return l
}

package simplematrix

import (
	"fmt"
)

type Matrix struct {
	elems [][]float64
	rows  uint
	cols  uint
}

func (m Matrix) String() string {
	return fmt.Sprintf("Matrix%v", m.elems)
}

func NewMatrix(rows, cols uint, e ...float64) (Matrix, error) {
	if rows*cols != uint(len(e)) {
		err := fmt.Errorf("colunas*linhas devem ser igual ao número de elementos da matriz")
		return Matrix{}, err
	}

	m := Matrix{}
	m.rows = rows
	m.cols = cols

	for r := 0; r < int(rows); r++ {
		line := make([]float64, cols)
		for c := 0; c < int(cols); c++ {
			line[c] = e[c+r*int(cols)]
		}
		m.elems = append(m.elems, line)
	}

	return m, nil
}

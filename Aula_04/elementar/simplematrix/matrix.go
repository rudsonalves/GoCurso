package simplematrix

import (
	"fmt"
)

type Matrix struct {
	elems [][]float64
	rows  uint
	cols  uint
}

func NewMatrix(rows, cols int, e ...float64) (Matrix, error) {
	if rows*cols != len(e) {
		err := fmt.Errorf("número de elementos (rows*cols) diferente de len(e)")
		return Matrix{}, err
	}

	m := Matrix{}
	for r := 0; r < rows; r++ {
		line := []float64{}
		for c := 0; c < cols; c++ {
			line = append(line, e[r*cols+c])
		}
		m.elems = append(m.elems, line)
	}

	m.rows = uint(rows)
	m.cols = uint(cols)

	return m, nil
}

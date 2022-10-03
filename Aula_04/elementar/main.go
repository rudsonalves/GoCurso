package main

import (
	"fmt"
	"rudsonalves/Leomar/Aula_04/elementar/statistics"
	"rudsonalves/Leomar/Aula_04/elementar_bkp/simplematrix"
)

func main() {
	statistics.Max(1, 2, 5, 4, 3)

	//statistics.Zero()

	m, err := simplematrix.NewMatrix(2, 3, 1, 2, 3, 4, 5, 6)

	fmt.Println(m, err)
}

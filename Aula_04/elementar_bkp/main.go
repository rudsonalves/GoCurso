package main

import (
	"fmt"
	"rudsonalves/Leomar/Aula_04/elementar/simplelist"
	"rudsonalves/Leomar/Aula_04/elementar/simplematrix"
	"rudsonalves/Leomar/Aula_04/elementar/statistics"
)

func main() {
	l0 := simplelist.NewList(1, 2, 3, 4)
	l1 := simplelist.NewList(5, 6, 7, 8)

	fmt.Println(statistics.MidRange(1, 2, 3, 4, 5, 6))

	fmt.Println(l0)
	fmt.Println(l0.Add(l1))

	m, _ := simplematrix.NewMatrix(2, 3, 1, 2, 3, 4, 5, 6)
	fmt.Println(m)
}

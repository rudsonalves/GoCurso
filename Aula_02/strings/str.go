package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := `Albert Einstein (Ulm, 14 de março de 1879 — Princeton, 18 de abril de 1955) foi um físico teórico alemão que desenvolveu a teoria da relatividade geral, um dos pilares da física moderna ao lado da mecânica quântica. Embora mais conhecido por sua fórmula de equivalência massa-energia, E = mc² — que foi chamada de "a equação mais famosa do mundo" —, foi laureado com o Prêmio Nobel de Física de 1921 "por suas contribuições à física teórica" e, especialmente, por sua descoberta da lei do efeito fotoelétrico, que foi fundamental no estabelecimento da teoria quântica.`

	contador := len(strings.Fields(frase))
	contt := strings.Count(frase, "—") + strings.Count(frase, "=")
	contador -= contt

	fmt.Println(contador)

	split := strings.Fields(frase)
	count := len(split)
	for _, p := range split {
		if strings.Contains(p, "—") || strings.Contains(p, "=") {
			count -= 1
		}
	}
	fmt.Println(count)
}

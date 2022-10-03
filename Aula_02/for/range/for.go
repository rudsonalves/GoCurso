package main

import "fmt"

func main() {
	estados := []string{"Amapa", "Acre", "Rondônia",
		"Rorâima", "Para", "Amazonas", "Mato Grosso do Sul",
		"Mato Grosso", "Goias", "Brasília", "Tocantins",
		"Minas Gerais", "Espírito Santo", "Rio de Janeiro",
		"São Paulo", "Rio Grande do Sul", "Santa Catarina",
		"Paraná", "Bahia", "Paraíba", "Sergipe", "Alagoas",
		"Maranhão", "Rio Grande do Norte", "Sergipe", "Ceará",
		"Pernanbuco",
	}

	for id, estado := range estados {
		fmt.Printf("%2d: %s\n", id+1, estado)
	}
	fmt.Println("")

	for _, estado := range estados {
		fmt.Printf("%s, ", estado)
	}
	fmt.Println("")
}

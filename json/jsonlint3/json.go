/*
Copyright ©2013 The jsontr Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/

/*
O aplicativo jsontr corrige json truncados. Este aplicativo é
corrige apenas json trucados removendo a última linha e alguns
caracteres indsejados como espaços e vírgula, d final de linha
anterior ao truncamento. Por fim o aplicativo adiciona os
fechamentos necessários (caracteres } e ]) para corrigir o json.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Position(input []byte, offset int) (line string, row, col int) {
	for _, c := range input[:offset] {
		if c == '\n' {
			row++
			col = 0
			line = ""
			continue
		} else if c == '\r' {
			continue
		}
		col++
		line += string(c)
	}

	return
}

// RemoveLastRow remove a última linha do json, e em seguida
// remove espaços e vírgula do final da última linha. Esta
// função percorre a slice input de trás para frente e remove
// a última linha e caracteres indesejados.
func RemoveLastRow(input []byte) []byte {
	id := len(input) - 2
	for id > 0 {
		if input[id] == '\n' {
			break
		}
		id--
	}
	id--

	for {
		if input[id] != ' ' && input[id] != ',' {
			break
		}
		id--
	}

	return input[:id+1]
}

// AddCloseKeys adiciona os caracteres } e ] necessarios para corrigir
// o json trucado. Esta função conta as aberturas e fechamentos dos
// { e [ para construir um slice de fechamentos correspondente.
func AddCloseKeys(input []byte) []byte {
	opened := []byte{}

	for _, c := range input {
		switch c {
		case '{':
			opened = append(opened, '}')
		case '}':
			if opened[len(opened)-1] == '}' {
				opened = opened[:len(opened)-1]
			} else {
				log.Fatal("Erro in }")
			}
		case '[':
			opened = append(opened, ']')
		case ']':
			if opened[len(opened)-1] == ']' {
				opened = opened[:len(opened)-1]
			} else {
				log.Fatal("Erro in ]")
			}
		default:
			continue
		}
	}

	output := input
	for id := len(opened) - 1; id >= 0; id-- {
		output = append(output, opened[id])
	}

	return output
}

func main() {
	fson, err := os.ReadFile("test2.json")
	if err != nil {
		log.Fatal(err)
	}

	var jsonInterface map[string]interface{}
	err = json.Unmarshal(fson, &jsonInterface)
	if err != nil {
		fson = RemoveLastRow(fson)
		fson = AddCloseKeys(fson)
		err = json.Unmarshal(fson, &jsonInterface)
		if err != nil {
			fmt.Println(string(fson))
			log.Fatal(err)
		}
	}
	err = json.Unmarshal(fson, &jsonInterface)
	if err != nil {
		log.Fatal(err)
	}

	behavior := jsonInterface["behavior"].(map[string]interface{})
	generic := behavior["generic"].([]interface{})
	generic1 := generic[1].(map[string]interface{})
	summary := generic1["summary"].(map[string]interface{})
	dll_loaded := summary["dll_loaded"].([]interface{})

	// indentOut, _ := json.MarshalIndent(dll_loaded, "", "\t")
	// fmt.Println(string(indentOut))

	for id, valor := range dll_loaded {
		fmt.Printf("[%02d] - %q\n", id, valor)
	}
}

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

func main() {
	fson, err := os.ReadFile("test2.json")
	if err != nil {
		log.Fatal(err)
	}

	fson = RemoveLastRow(fson)

	var jsonInterface interface{}
	char := byte('}')
	for {
		err = json.Unmarshal(fson, &jsonInterface)
		if err == nil {
			break
		}

		fson = append(fson, char)
		switch e := err.(type) {
		case *json.SyntaxError:
			if e.Error() != "unexpected end of JSON input" {
				char = byte(']')
				fson = fson[:len(fson)-2]
			} else {
				char = byte('}')
			}
		default:
			log.Fatal(e.Error())
		}
	}

	indentOut, err := json.MarshalIndent(jsonInterface, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(indentOut))
}

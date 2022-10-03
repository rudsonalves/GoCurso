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
	fson = append(fson, ']', '}', '}', ']', '}', '}')

	var jsonInterface interface{}
	err = json.Unmarshal(fson, &jsonInterface)
	if err != nil {
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			line, row, col := Position(fson, int(e.Offset))
			fmt.Printf("UnmarshalTypeError: %s - Value[%s]\n", e.Error(), e.Value)
			fmt.Printf(" Type[%v] offset: %d row: %d  col: %d\n", e.Type, e.Offset, row, col)
			fmt.Printf("line> %s<<\n", line)
		case *json.SyntaxError:
			line, row, col := Position(fson, int(e.Offset))
			fmt.Printf("SyntaxError: %s\n", e.Error())
			fmt.Printf(" Offset: %d row: %d, col: %d\n", e.Offset, row, col)
			fmt.Printf("line>> %s<<\n", line)
		case *json.InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: %v, Type[%v]\n", e.Error(), e.Type)
		default:
			fmt.Printf("error: %v", e.Error())
		}
	}

	_, err = json.MarshalIndent(jsonInterface, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(indentOut))
}

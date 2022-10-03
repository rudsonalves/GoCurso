package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}]`)

	type Animal map[string]interface{}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}

	a := animals[0]
	fmt.Printf("%v %[1]T\n", a["Name"])

	b := map[string]string{
		"Name":  "Platypus",
		"Order": "Monotremata",
	}

	fmt.Printf("%v %[1]T\n", b)
}

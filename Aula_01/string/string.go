package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := strconv.Itoa(45)
	b, _ := strconv.Atoi("45")
	fmt.Printf("%v %[1]T\n", a)
	fmt.Printf("%v %[1]T\n", b)

	c, _ := strconv.ParseBool("1")
	fmt.Println(c)

	d := "Ação e reação"
	fmt.Println(strings.Contains(d, "ão"))
	fmt.Println(strings.Count(d, "ão"))
	fmt.Println(strings.ReplaceAll(d, "e", "E"))
	fmt.Println(strings.Index(d, "ão"))
	fmt.Println(strconv.Quote(d))
}

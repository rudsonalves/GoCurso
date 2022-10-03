package main

import "fmt"

func main() {
	// var x interface{}
	// var x interface{} = 5.
	// var x interface{} = 5
	// var x interface{} = "5"
	var x interface{} = 2 + 3i

	if v, ok := x.(float64); ok {
		fmt.Printf("a is a float64 and you value is: %.2f\n", v)
	}

	if v, ok := x.(string); ok {
		fmt.Printf("a is a string and you value is: %s\n", v)
	}

	if v, ok := x.(int); ok {
		fmt.Printf("a is a string and you value is: %d\n", v)
	}

	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x is int and you value is", v)
	case float64:
		fmt.Println("x is float64 and you value is", v)
	case string:
		fmt.Println("x is string and you value is", v)
	default:
		fmt.Println("don't know the type and you value is", v)
	}
}

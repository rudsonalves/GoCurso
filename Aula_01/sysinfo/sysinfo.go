package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println(runtime.Compiler, runtime.GOARCH, runtime.GOOS)
	fmt.Println(strconv.IntSize)
}

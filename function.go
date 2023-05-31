package main

import (
	"fmt"
)

func main() {
	first := 5
	second := 3
	fmt.Println(add(first, second))
}

func add(a int, b int) int {
	return a + b
}

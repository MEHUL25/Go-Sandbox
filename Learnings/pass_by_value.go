package main

import (
	"fmt"
)

func main() {
	x := 4
	y := x
	fmt.Printf("x= %v y= %v", x, y)

	x = 6
	fmt.Printf("x= %v y= %v", x, y)
}

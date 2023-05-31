package main

import (
	"fmt"
)

func unnamed_return() (int, int) { // unnamed or exploicit returns
	x := 2
	y := 3
	return x, y
}

func named_return() (a int, b int) { // named return - implicit returns
	a = 5
	b = 6
	return
}

func main() {

	x, y := unnamed_return()
	fmt.Printf("x= %v y = %v", x, y)

	p, _ := unnamed_return()
	fmt.Printf("p= %v ", p)

	a, b := named_return()
	fmt.Printf("a= %v b=%v ", a, b)

}

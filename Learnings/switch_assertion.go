package main

import (
	"fmt"
)

func test(obj interface{}) {
	switch v := obj.(type) {
	case int:
		fmt.Println("%T", v)
	case string:
		fmt.Println("%T", v)
	default:
		fmt.Println("default case : %T", v)
	}
}

func main() {
	test(12)
	test(23.45)
	test("It is a string")
}

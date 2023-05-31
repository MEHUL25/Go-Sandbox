package main

import (
	"fmt"
)

func populate() [5]int {
	var array [5]int
	for i := 0; i < 5; i++ {
		array[i] = i
	}
	return array
}

func display(array [5]int) {
	l := len(array)
	for i := 0; i < l; i++ {
		fmt.Printf("%v\n", array[i])
	}
	fmt.Println("Array Dsiplayed successfully")
}

func main() {
	var array [5]int

	array = populate()

	display(array)
}

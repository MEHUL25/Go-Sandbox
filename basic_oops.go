package main

import (
	"fmt"
)

type rectangle struct {
	leng int
	bred int
}

func (object rectangle) area() int {
	return object.leng * object.bred
}

func main() {
	r := rectangle{
		leng: 5,
		bred: 10,
	}

	fmt.Println(r.area())
}

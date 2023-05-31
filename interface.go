package main

import (
	"fmt"
)

type shape interface {
	area() int
	perimeter() int
}

type rectangle struct {
	width, height int
}

type circle struct {
	radius int
}

func (r rectangle) area() int {
	return r.width * r.height
}

func (r rectangle) perimeter() int {
	return r.height + r.width
}

func (c circle) area() int {
	area := int(3 * c.radius * c.radius)
	return area
}

func (c circle) perimeter() int {
	return int(2 * 3 * c.radius)
}

func calculate(object shape) {
	fmt.Println("Object ", object)
	fmt.Println("Area = ", object.area())
	fmt.Println("perimeter = ", object.perimeter())
}

func main() {
	r := rectangle{
		height: 5,
		width:  10,
	}
	calculate(r)
	fmt.Println("Circle")
	fmt.Println()
	c := circle{
		radius: 5,
	}
	calculate(c)
}

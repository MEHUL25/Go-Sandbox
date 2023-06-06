package main

import (
	"fmt"
)

func populate(row, column int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < row; i++ {
		singlerow := make([]int, 0)
		for j := 0; j < column; j++ {
			singlerow = append(singlerow, i*j)
		}
		matrix = append(matrix, singlerow)
	}
	return matrix
}

func display(matrix [][]int) {
	r := len(matrix)
	c := len(matrix[0])
	fmt.Printf("r=%d c=%d\n\n", r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func test(row, column int) {
	matrix := populate(row, column)
	display(matrix)
}

func main() {
	test(3, 4)
}

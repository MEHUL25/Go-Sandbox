package main

import "fmt"

type custError struct {
	err int
}

func (e custError) Error() string {
	return fmt.Sprintf("Error %v", e.err)
}

func check(dividend int) (int, error) {
	if dividend == 0 {
		return 0, custError{err: 0}
	}
	return dividend, nil

}

func test(dividend int) {
	result, err := check(dividend)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func main() {
	test(0)
	test(2)
}

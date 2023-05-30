package main

import (
	"fmt"
)

func main() {
	height := 5
	if height < 5 {
		fmt.Println("Short")
	} else if height >= 5 && height < 10 {
		fmt.Println("Medium")
	} else {
		fmt.Println("Tall")
	}
}

package main

import (
	"fmt"
)

func populate(m map[string]int, mob []int, name [5]string) {
	l := len(mob)
	for i := 0; i < l; i++ {
		m[name[i]] = mob[i]
	}
}

func main() {
	m := make(map[string]int)

	mob := make([]int, 0)
	name := [5]string{"fsv", "gegveg", "eveve", "dvdvdv", "fvfdvdf"}

	for i := 0; i < 5; i++ {
		mob = append(mob, i)
	}

	populate(m, mob, name)
	fmt.Println(m)

	for i, j := range m {
		fmt.Printf("%v %v \n", i, j)
	}

	delete(m, "fsv")

	fmt.Println(m)

}
